package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"runtime"
	"strings"
	"time"

	"github.com/henrikig/feature-selection/regression"
	"github.com/sjwhitworth/golearn/base"
)

type Work struct {
	bitString string
	data      *base.DenseInstances
	nonClass  []base.Attribute
	class     []base.Attribute
	rmse      float64
}

func worker(in <-chan *Work, out chan<- *Work) {
	defer close(out)

	for w := range in {
		newInst := regression.GetColumns(w.data, w.bitString, w.nonClass, w.class)
		w.rmse = regression.GetFitness(newInst)
		fmt.Println("Calculated rmse")
		out <- w
		fmt.Println("Sent to out channel")
	}
}

func main() {
	nWorkers := runtime.NumCPU() - 1
	run(nWorkers)
}

func run(nWorkers int) {
	data, err := base.ParseCSVToInstances("./datasets/dataset.csv", false)

	if err != nil {
		log.Fatal(err)
	}

	//data.AddClassAttribute(base.NewFloatAttribute("price"))
	// data.RemoveClassAttribute(base.NewFloatAttribute("sqft_lot15"))

	n := len(base.NonClassAttributes(data))

	nonClass := base.NonClassAttributes(data)
	class := data.AllClassAttributes()

	start := time.Now()

	in, out := make(chan *Work, nWorkers), make(chan *Work, nWorkers)

	for i := 0; i < nWorkers; i++ {
		go worker(in, out)
	}

	go func() {
		for i := 0; i < 100; i++ {
			work := &Work{
				bitString: generateBitString(n),
				data:      data,
				nonClass:  nonClass,
				class:     class,
			}

			in <- work
			fmt.Println("Sent", i+1)
		}
		close(in)
	}()

	var bestWork = &Work{
		rmse: math.Inf(1),
	}

	for w := range out {
		fmt.Println(w.bitString, w.rmse)

		if w.rmse < bestWork.rmse {
			bestWork = w
		}
	}

	fmt.Println("The best bitstring and rmse are:", bestWork.bitString, "&", bestWork.rmse)
	elapsed := time.Since(start)
	fmt.Println("Process took", elapsed)
}

func generateBitString(n int) string {
	builder := strings.Builder{}

	for i := 0; i < n; i++ {
		num := rand.Intn(2)
		if num == 0 {
			builder.WriteString("0")
		} else {
			builder.WriteString("1")
		}
	}

	return builder.String()
}

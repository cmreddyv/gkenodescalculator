//get maximum number of pods per node

//number of pods allocated to every node = 2 * maximum pods per node = 2 * 32 = 64 = 2^6

//number of ips in the cluster pod range = 10.96.2.0/23 = 2 ^(32- pod cir range) = 2 ^ 9

// maximum number od node = 2 ^ 9 divided by 2 ^ 6 = 2 ^ 3 = 8 nodes

package main

import (
	"fmt"
	"math"
)

func main() {

	var i, j int

	fmt.Print("enter the max number of pods in node pool: ")

	fmt.Scan(&i)

	fmt.Print("enter the subnet range of the pods between 8 and 32: ")

	fmt.Scan(&j)

	fmt.Println("j is:", j)

	if j < 8 || j > 32 {

		fmt.Println("enter the range of pod subnet between 8 and 32 only")

	} else {

		k := float64((32 - j))

		fmt.Println("maximum number of node in a cluster: ", int64(math.Pow(2, k)/float64(2*i)))
	}

}

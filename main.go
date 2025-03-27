//get maximum number of pods per node

//number of pods allocated to every node = 2 * maximum pods per node = 2 * 32 = 64 = 2^6

//number of ips in the cluster pod range = 10.96.2.0/23 = 2 ^(32- pod cir range) = 2 ^ 9

// maximum number od node = 2 ^ 9 divided by 2 ^ 6 = 2 ^ 3 = 8 nodes

package main

import (
	"fmt"
	"log"
	"math"
	"os"
)

// func main() {

// 	var i, j int

// 	fmt.Print("enter the max number of pods in node pool: ")

// 	fmt.Scan(&i)

// 	fmt.Print("enter the subnet range of the pods between 8 and 32: ")

// 	fmt.Scan(&j)

// 	if j < 8 || j > 32 {

// 		fmt.Println("enter the range of pod subnet between 8 and 32 only")

// 	} else {

// 		k := float64((32 - j))

// 		fmt.Println("maximum number of node in a cluster: ", int64(math.Pow(2, k)/float64(2*i)))
// 	}

// }

type PodClusterCalculator struct {
	MaxPodsPerNode int
	SubnetRange    int
}

func (c *PodClusterCalculator) Validate() error {

	if c.MaxPodsPerNode <= 0 {
		return fmt.Errorf("max pods per node should be a positive number ")
	}

	if c.SubnetRange < 8 || c.SubnetRange > 32 {
		return fmt.Errorf("pods subnet range should be in between 8 and 32")
	}

	return nil

}

func (c *PodClusterCalculator) CalculateMaxNodes() int {
	availableIPs := math.Pow(2, float64(32-c.SubnetRange))
	maxNodesInCluster := int(availableIPs / float64(2*c.MaxPodsPerNode))
	return maxNodesInCluster
}

func (c *PodClusterCalculator) PrintClusterInfo() {
	fmt.Printf("Cluster Configuration Details:\n")
	fmt.Printf("- Max Pods Per Node: %d\n", c.MaxPodsPerNode)
	fmt.Printf("- Pod Subnet Range: /%d\n", c.SubnetRange)
	fmt.Printf("- Maximum Nodes in Cluster: %d\n", c.CalculateMaxNodes())
}

func main() {
	calculator := &PodClusterCalculator{}

	fmt.Print("enter the max number of pods in Nodepool: ")

	_, err := fmt.Scan(&calculator.MaxPodsPerNode)

	if err != nil {
		log.Fatal("Invalid input number of nodes")
	}

	fmt.Print("Enter the pods subnet range (between 8 and 32): ")

	_, err = fmt.Scan(&calculator.SubnetRange)

	if err != nil {
		log.Fatal("Invalid input for subnet range")
	}

	if err := calculator.Validate(); err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	calculator.PrintClusterInfo()

}

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
	"strconv"
	"strings"
)

type PodClusterCalculator struct {
	MaxPodsPerNode int
	SubnetRange    int
}

func ExtractSubnetMask(cidrNotation string) (int, error) {

	parts := strings.Split(cidrNotation, "/")

	if len(parts) != 2 {
		return 0, fmt.Errorf("Invalid CIDR notation: %s", cidrNotation)
	}

	subnetMask, err := strconv.Atoi(parts[1])

	if err != nil {
		return 0, fmt.Errorf("invalid subnet mask: %s", parts[1])
	}

	if subnetMask < 8 || subnetMask > 32 {
		return 0, fmt.Errorf("subnet mask must be between 8 and 32, got: %d", subnetMask)
	}

	return subnetMask, nil
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
	//fmt.Printf("- Pod Subnet Range: /%d\n", c.SubnetRange)
	fmt.Printf("- Maximum Nodes in Cluster: %d\n", c.CalculateMaxNodes())
}

func main() {
	calculator := &PodClusterCalculator{}

	fmt.Print("enter the max number of pods in Nodepool: ")

	_, err := fmt.Scan(&calculator.MaxPodsPerNode)

	if err != nil {
		log.Fatal("Invalid input number of nodes")
	}

	var cidrNotation string

	fmt.Print("Enter the pod subnet CIDR (e.g., 10.0.0.0/23): ")

	_, err = fmt.Scan(&cidrNotation)

	if err != nil {
		log.Fatal("Invalid input for CIDR notation")
	}

	calculator.SubnetRange, err = ExtractSubnetMask(cidrNotation)

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	if err := calculator.Validate(); err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	calculator.PrintClusterInfo()

}

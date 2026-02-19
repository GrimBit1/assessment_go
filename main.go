package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("cities.csv")
	if err != nil {
		log.Fatalf("file read errored:%v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	for {
		record, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf("Error reading CSV record: %s", err)
		}
		cityCode := strings.TrimSpace(record[0])
		if cityCode == "City Code" {
			continue
		}
		// Trim spaces
		provinceCode := strings.TrimSpace(record[1])
		countryCode := strings.TrimSpace(record[2])

		cityName := strings.TrimSpace(record[3])
		provinceName := strings.TrimSpace(record[4])
		countryName := strings.TrimSpace(record[5])

		// Add it to map for indexing
		if _, ok := regionMap[countryCode]; !ok {
			regionMap[countryCode] = NewRegion(countryName, World)
		}

		// Add it to map for indexing
		if _, ok := regionMap[provinceCode+"-"+countryCode]; !ok {
			regionMap[provinceCode+"-"+countryCode] = NewRegion(provinceName, regionMap[countryCode])
		}

		// Add it to map for indexing
		if _, ok := regionMap[cityCode+"-"+provinceCode+"-"+countryCode]; !ok {
			regionMap[cityCode+"-"+provinceCode+"-"+countryCode] = NewRegion(cityName, regionMap[provinceCode+"-"+countryCode])
		}

	}

	if len(regionMap) == 0 {
		log.Fatalf("no rows in csv")
	}

	// DISTRIBUTOR1
	d1 := NewDistributor("DISTRIBUTOR1", nil)
	err = d1.AddInclude("IN", "US")
	if err != nil {
		fmt.Println("Can't exclude IN,US:", err)
	}
	err = d1.AddExclude("KA-IN")
	if err != nil {
		fmt.Println("Can't exclude KA-IN:", err)
	}
	err = d1.AddExclude("CENAI-TN-IN")
	if err != nil {
		fmt.Println("Can't exclude CENAI-TN-IN:", err)
	}

	fmt.Println("\n----- DISTRIBUTOR1 -----")
	fmt.Println("Chicago:", d1.HasPermission("CHIAO-IL-US"))
	fmt.Println("Chennai:", d1.HasPermission("CENAI-TN-IN"))
	fmt.Println("Vadodara:", d1.HasPermission("VODRA-GJ-IN"))
	fmt.Println("Karnataka:", d1.HasPermission("KA-IN"))
	fmt.Println("Sagar:", d1.HasPermission("SAGAR-KA-IN"))

	// DISTRIBUTOR2 < DISTRIBUTOR1
	d2 := NewDistributor("DISTRIBUTOR2", d1)
	err = d2.AddInclude("IN")
	if err != nil {
		fmt.Println("Can't include IN:", err)
	}
	err = d2.AddExclude("TN-IN")
	if err != nil {
		fmt.Println("Can't exclude TN-IN:", err)
	}

	fmt.Println("\n----- DISTRIBUTOR2 -----")
	fmt.Println("Hublingen:", d2.HasPermission("HUBLE-RP-DE"))
	fmt.Println("Tamil Nadu:", d2.HasPermission("TN-IN"))
	fmt.Println("Maharashtra:", d2.HasPermission("MH-IN"))

	d3 := NewDistributor("DISTRIBUTOR3", d2)

	// DISTRIBUTOR3 is authorized to distribute only Vadodara, Gujarat, India.
	if err := d3.AddInclude("VODRA-GJ-IN"); err != nil {
		fmt.Println("Can't include VODRA-GJ-IN:", err)
	}

	fmt.Println("\n----- DISTRIBUTOR3 -----")
	fmt.Println("Vadodara:", d3.HasPermission("VODRA-GJ-IN"))
	fmt.Println("Mumbai:", d3.HasPermission("MUBAI-MH-IN"))
	fmt.Println("India:", d3.HasPermission("IN"))
}

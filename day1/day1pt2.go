package main

import (
    "bufio"
	"fmt"
	"log"
	"os"
    "strconv"
)

func main() {

    file, err := os.Open("inputs.txt")
    if err != nil {
        log.Fatal(err)
    }
 
    var measurements []int

    scanner := bufio.NewScanner(file)

	for scanner.Scan() {
        measurement, err := strconv.Atoi(scanner.Text())
        if err != nil {
            log.Fatal(err)
        }

        measurements = append(measurements, measurement)
    }

    // got the array, now itirate through edge values since the values inbetween are common between the two
    count := 0
    for i := 0; i <= len(measurements)-4; i++ {
        if measurements[i] < measurements[i+3]{
            count++
        } 
    }

    fmt.Printf("Count: %02d", count);

    if scanner.Err() != nil {
		log.Println(scanner.Err())
	}
}
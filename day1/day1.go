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
   
    scanner := bufio.NewScanner(file)

	counter := -1;
    previousMeasurement := 0;
	for scanner.Scan() {
        depthMeasurement, err := strconv.Atoi(scanner.Text())
        if err != nil {
            log.Fatal(err)
        }

        if depthMeasurement > previousMeasurement {
          counter++
        };
        
        previousMeasurement = depthMeasurement; 
	}

    fmt.Printf("Count: %02d", counter);

    if scanner.Err() != nil {
		log.Println(scanner.Err())
	}
}
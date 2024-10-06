# Parallel Sum of JSON Data

## Description

This Go program reads a JSON file containing an array of objects, where each object has two integer fields \(a\) and \(b\) with values in the range \([-10, 10]\). The program calculates the sum of all these numbers in parallel using goroutines for efficient processing. The total sum is then printed to the console.

## Features

- Reads a JSON file with 1,000,000 objects.
- Each object contains two integer fields, \(a\) and \(b\).
- Uses goroutines to perform calculations in parallel.
- Allows the user to specify the number of goroutines to use for computation.

## Requirements

- Go 1.16 or later
- A JSON file containing the data (see [large_data.json])

## Getting Started

1. **Clone the Repository**

   ```bash
   git clone <repository-url>
   cd <repository-directory>

2. **Run the program and add number of goroutines to the arguments**

   ```bash
   go run main.go 2
   

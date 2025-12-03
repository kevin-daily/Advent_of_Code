package main

func splitNumber (numToSplit string) ([]int) {
	// TODO: Split number in half
	// If length is odd, ignore. 
	// Find half of length and split into two parts


}

func checkDoubling (number int) (bool) {
	var isDouble = False

	var numbers []int
	numbers := splitNumber(number)

	if (numbers[1]==numbers[2]) {
		isDouble := True
	}
	
	return isDouble
}

func main() {
	// TODO: Read in text from file
	file, er := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		input := scanner.Text()
	}
	var rangeStart, rangeEnd int

	for i := rangeStart; i <= rangeEnd; i++ {

	} 
}
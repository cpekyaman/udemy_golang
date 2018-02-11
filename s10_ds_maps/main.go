package main

import (
	"bufio"
	"log"
	"net/http"
	"fmt"
)

func Hash(sz int, s string) int {
	sum := 0
	for _, c := range s {
		sum += int(c)
	}
	return sum % sz
}

func hash_table(bucketSize int) [][]string {
	res, err := http.Get("http://www.gutenberg.org/files/2701/2701-0.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close()

	scanner.Split(bufio.ScanWords)
	
	buckets := make([][]string, bucketSize)
	for i := 0; i < bucketSize; i++ {
		buckets = append(buckets, []string{})
	}

	for scanner.Scan() {
		w := scanner.Text()
		n := Hash(bucketSize, w)
		buckets[n] = append(buckets[n], w)
	}

	return buckets
}

func main() {
	m := make(map[string]int)
	m["k1"] = 10
	m["k2"] = 12

	fmt.Println(m)

	_, ok := m["k2"]
	fmt.Println(ok)
	delete(m, "k2")

	_, ok = m["k2"]
	fmt.Println(ok)

	m["k2"] = 15
	m["k3"] = 13
	m["k4"] = 14

	for k, v := range m {
		fmt.Println(k, "-", v)
	}

	const bucketSize = 20
	ht := hash_table(bucketSize)

	for i := 0; i < bucketSize; i++ {
		fmt.Println("len(", i , ") :", len(ht[i]))
	}
}
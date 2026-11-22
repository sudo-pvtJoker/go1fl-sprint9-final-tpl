package main

import (
	"log"
	"math"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	if size <= 0 {
		return []int{}
	}

	data := make([]int, size)
	for i := range data {
		data[i] = rand.Intn(math.MaxInt) + 1
	}
	return data
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}

	largest := data[0]
	for _, value := range data[1:] {
		if value > largest {
			largest = value
		}
	}
	return largest
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}

	parts := min(CHUNKS, len(data))
	chunkSize := len(data) / parts
	maxima := make([]int, parts)
	var wg sync.WaitGroup
	wg.Add(parts)
	for i := 0; i < parts; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if i == parts-1 {
			end = len(data)
		}
		go func(index int, chunk []int) {
			defer wg.Done()
			maxima[index] = maximum(chunk)
		}(i, data[start:end])
	}
	wg.Wait()
	return maximum(maxima)
}

func main() {
	log.Printf("Генерируем %d целых чисел", SIZE)
	data := generateRandomElements(SIZE)

	log.Printf("Ищем максимальное значение в один поток")
	start := time.Now()
	result := maximum(data)
	elapsed := time.Since(start).Milliseconds()
	log.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms", result, elapsed)

	log.Printf("Ищем максимальное значение в %d потоков", CHUNKS)
	start = time.Now()
	result = maxChunks(data)
	elapsed = time.Since(start).Milliseconds()
	log.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms", result, elapsed)
}

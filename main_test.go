package main

import "testing"

func TestGenerateRandomElements(t *testing.T) {
	for _, size := range []int{-1, 0, 1, 100} {
		data := generateRandomElements(size)
		if len(data) != max(0, size) {
			t.Errorf("Для размера %d получили слайс длины %d", size, len(data))
		}
		for _, value := range data {
			if value <= 0 {
				t.Errorf("Для размера %d получили неположительное число %d", size, value)
			}
		}
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{"nil", nil, 0},
		{"пустой слайс", []int{}, 0},
		{"один элемент", []int{7}, 7},
		{"максимум в начале", []int{9, 3, 1}, 9},
		{"максимум в конце", []int{1, 3, 9}, 9},
		{"повтор максимума", []int{5, 2, 5}, 5},
		{"отрицательные числа", []int{-4, -1, -3}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximum(tt.data); got != tt.want {
				t.Errorf("maximum() = %d, ожидалось %d", got, tt.want)
			}
		})
	}
}

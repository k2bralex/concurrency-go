package md_1

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*Задание 1
Создайте консольное приложение, порождающее поток.
Этот поток должен отображать в консоль числа от 0 до n. */

/*Задание 2
Добавьте в первое задание возможность передачи начала и конца диапазона чисел.
Границы определяет пользователь.*/

/*Задание 3
Добавьте к первому заданию возможность определения пользователем количества потоков.
Границы диапазона чисел также выбираются пользователем.*/

func PrintRange(from, too, gor int) {
	var wg sync.WaitGroup

	wg.Add(gor)
	for i := 0; i < gor; i++ {
		go func() {
			for i := from; i <= too; i++ {
				fmt.Printf("%d ", i)
				time.Sleep(10 * time.Millisecond)
			}
			wg.Done()
		}()
	}

	wg.Wait()

}

/*Задание 4
Консольное приложение генерирует набор чисел, состоящий из 10000 элементов.
С помощью механизма потоков нужно найти максимум, минимум, среднее в этом наборе.
Для каждой из задач выделите поток.*/

/*Задание 5
К четвертому заданию добавьте поток, выводящий набор чисел и результаты вычислений в файл.*/

func MinMaxAvgGoroutines(min, max, len int) {

	slPtr := fillSlice(min, max, len)

	start := time.Now()

	minVal := make(chan int)
	maxVal := make(chan int)
	avgVal := make(chan int64)

	findMin(max, slPtr, minVal)
	findMax(min, slPtr, maxVal)
	findAvg(slPtr, avgVal)

	fmt.Printf(
		"In range %d : %d, len = %d\nMin = %d\nMax = %d\nAvg = %d\n",
		min,
		max,
		len,
		<-minVal,
		<-maxVal,
		<-avgVal,
	)

	fmt.Println(time.Now().Sub(start).Nanoseconds())
}

func MinMaxAvgNoGoroutines(min, max, len int) {

	slPtr := fillSlice(min, max, len)

	start := time.Now()

	mi := findMinLine(max, slPtr)
	ma := findMaxLine(min, slPtr)
	av := findAvgLine(slPtr)

	fmt.Printf(
		"In range %d : %d, len = %d\nMin = %d\nMax = %d\nAvg = %d\n",
		min,
		max,
		len,
		mi,
		ma,
		av,
	)

	fmt.Println(time.Now().Sub(start).Nanoseconds())
}

func findAvgLine(s *[]int) int64 {
	var sum int64 = 0
	for _, v := range *s {
		sum += int64(v)
	}
	return sum / int64(len(*s))
}

func findMaxLine(min int, s *[]int) int {
	max := min
	for _, v := range *s {
		if v > max {
			max = v
		}
	}
	return max

}

func findMinLine(max int, s *[]int) int {
	min := max
	for _, v := range *s {
		if v < min {
			min = v
		}
	}
	return min
}

func findAvg(s *[]int, ch chan int64) chan int64 {
	go func() {
		var sum int64 = 0
		for _, v := range *s {
			sum += int64(v)
		}
		ch <- sum / int64(len(*s))
		close(ch)
	}()
	return ch
}

func findMax(min int, s *[]int, ch chan int) chan int {
	go func() {
		max := min
		for _, v := range *s {
			if v > max {
				max = v
			}
		}
		ch <- max
		close(ch)
	}()
	return ch
}

func findMin(max int, s *[]int, ch chan int) chan int {
	go func() {
		min := max
		for _, v := range *s {
			if v < min {
				min = v
			}
		}
		ch <- min
		close(ch)
	}()
	return ch
}

func fillSlice(min, max, len int) *[]int {
	var slice []int
	for i := 0; i < len; i++ {
		slice = append(slice, randomGen(min, max))
	}
	return &slice
}

func fillSliceGoroutines(min, max, len int) *[]int {
	var slice []int

	for i := 0; i < len; i++ {
		slice = append(slice, randomGen(min, max))
	}
	return &slice
}

func randomGen(min, max int) int {
	rand.NewSource(time.Now().UnixNano())
	return rand.Intn(max-min) + min + 1

}

func finMaxDivisior(n int) int {
	num := 2
	for n%num == 0 {
		num++
	}
	return n / num
}

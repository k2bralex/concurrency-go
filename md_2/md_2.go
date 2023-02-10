package md_2

import (
	"fmt"
	"strings"
	"sync"
	"unicode"
)

/*Задание 1
Создайте приложение, которое позволяет пользователю запустить несколько потоков.
Один поток генерирует числа, другой буквы, третий символы.*/

func LineGenerator() {
	var wg sync.WaitGroup
	wg.Add(3)

	go digitGen(&wg)
	go letterGen(&wg)
	go symbGen(&wg)

	wg.Wait()
}

func digitGen(wg *sync.WaitGroup) {
	for i := 0; i < 100; i++ {
		fmt.Println(i)
	}
	wg.Done()
}

func letterGen(wg *sync.WaitGroup) {
	k := 0
	for k < 100 {
		for i := 33; i < 127; i++ {
			r := []rune(string(i))[0]
			if unicode.IsLetter(r) {
				fmt.Println("     ", string(i))
				k++
			}
		}
	}
	wg.Done()
}

func symbGen(wg *sync.WaitGroup) {
	k := 0
	for k < 100 {
		for i := 33; i < 127; i++ {
			r := []rune(string(i))[0]
			if unicode.IsSymbol(r) {
				fmt.Println("          ", string(i))
				k++
			}
		}
	}
	wg.Done()
}

/*Задание 2
Создайте оконное приложение для подсчета факториала числа. Пользователь вводит число.
Приложение подсчитывает факториал этого числа. Для подсчета степени используйте механизм асинхронности.*/

/*Задание 3
Приложение для подсчета степени числа. Пользователь вводит число и степень.
Приложение подсчитывает степень числа. Для подсчета степени используйте механизм асинхронности.*/

func FactorPower(num, pow int) {
	f := factor(num)
	p := power(num, pow)

	for v := range f {
		fmt.Println(v)
	}
	for v := range p {
		fmt.Println(v)
	}
}

func factor(num int) chan int {
	out := make(chan int)
	go func() {
		result := 1
		for num > 0 {
			result *= num
			num--
		}
		out <- result
		close(out)
	}()
	return out
}

func power(num, pow int) chan int {
	out := make(chan int)
	go func() {
		for i := 1; i <= pow; i++ {
			num *= pow
		}
		out <- num
		close(out)
	}()
	return out
}

/*Задание 4
Пользователь вводит некоторый текст. Приложение подсчитывает количество гласных, согласных, символов.
Используйте механизм асинхронности для реализации задачи.*/

func Report(s string) {
	v := vowelsCount(&s)
	c := consonantsCount(&s)
	sy := symbolsCount(&s)

	fmt.Printf("vowels: %d, consonant: %d, symbols: %d", <-v, <-c, <-sy)
}

func vowelsCount(s *string) chan int {
	out := make(chan int)
	go func() {
		vowels := []rune("euioa")
		s1 := strings.ToLower(*s)
		count := 0
		for _, v := range vowels {
			count += strings.Count(s1, string(v))
		}
		out <- count
		close(out)
	}()
	return out
}

func consonantsCount(s *string) chan int {
	out := make(chan int)
	go func() {
		vowels := []rune("qwrtypsdfghjklzxcvbnm")
		s1 := strings.ToLower(*s)
		count := 0
		for _, v := range vowels {
			count += strings.Count(s1, string(v))
		}
		out <- count
		close(out)
	}()
	return out
}

func symbolsCount(s *string) chan int {
	out := make(chan int)
	go func() {
		s1 := strings.ToLower(*s)
		count := 0
		for v := range []rune(s1) {
			if unicode.IsSymbol(rune(v)) {
				count++
			}
		}
		out <- count
		close(out)
	}()
	return out
}

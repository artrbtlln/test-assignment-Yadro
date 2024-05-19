package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)      // Создаем reader для чтения ввода
	fmt.Print("Введите число контейнеров: ") // Ввод
	input, _ := reader.ReadString('\n')      // Считываем ввод
	input = strings.TrimSpace(input)         // Удаляем лишние пробелы и переводы строк
	n, _ := strconv.Atoi(input)              //Преобразуем строку в int

	balls := make([][]int, n)                        // Создаем двумерный срез размера "n*n"
	fmt.Println("Распределите шары по контейнерам:") // Ввод
	for i := 0; i < n; i++ {
		cont, _ := reader.ReadString('\n') // Считываем ввод
		cont = strings.TrimSpace(cont)     //Удаляем лишние пробелы и переводы строк
		nums := strings.Split(cont, " ")   //Разделяем строку
		balls[i] = make([]int, n)
		for j, str := range nums {
			balls[i][j], _ = strconv.Atoi(str) // Проходимся циклом, преобразовывая каждую часть в int и сохраняем в срез
		}
	}
	// Проверка возможности сортировки
	if AbilityToSort(n, balls) {
		fmt.Println("Можно отсортировать")
	} else {
		fmt.Println("Нельзя отсортировать")
	}
}
func AbilityToSort(n int, balls [][]int) bool { // Создаем функцию принимающую размер и срез
	colorSum := make([]int, n) // Сумма шаров по цветам
	ballsSum := make([]int, n) // Сумма шаров в контейнерах

	// Проходимся циклом и считаем суммы
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			ballsSum[i] += balls[i][j]
			colorSum[j] += balls[i][j]
		}
	}
	// Сортировка и ее варианты
	insertionSort(ballsSum)
	insertionSort(colorSum)
	//BubbleSort(ballsSum)
	//BubbleSort(colorSum)
	//sort.Ints(ballsSum)
	//sort.Ints(colorSum)

	//Проверяем равенство массивов
	for i := 0; i < n; i++ {
		if ballsSum[i] != colorSum[i] {
			return false
		}
	}
	return true
}
func insertionSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := i; j > 0 && arr[j-1] > arr[j]; j-- {
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}
	return arr
}
func BubbleSort(arr []int) []int {
	size := len(arr)
	for i := 0; i < size-1; i++ {
		for j := 0; j < size-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr

}

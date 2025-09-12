package main

func someFunc() string {
	v := createHugeString(1 &lt;&lt; 10)
	result := make([]byte, 100)
	copy(result, v[:100])
	return string(result)
}

func main() {
	justString := someFunc()
}

// Проблема была в том что juststring ссылалалась на всю сгенерированную строку, тем самым занимала 1024 байта памяти
// Теперь функция возвращает копию на 100 символов от полученной строки в локальную переменную

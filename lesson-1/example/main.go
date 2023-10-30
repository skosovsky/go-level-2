package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

/*
Задача: в заданной директории прочитать все файлы и сказать сколько раз встречается строка "Hello Word" в файлах и в каких

Hints:
	* для запуска `go run main.go -dirPath <путь до директории>`
	* что бы вызвать recover() из 49 строки - передать в качестве аргумента не существующую директорию
	* что бы увидеть "обернутую" ошибку в hasSubstring - передать в качестве аргумента не существующую директорию
	* что бы увидеть "самописную" ошибку из getStats забери у файла из директории права на чтение `chmod -r <имя_файла>` (+r вернет права)
*/

var (
	// говорим, что в переменную dirPath нужно будет положить значение
	// dirPath имеет тип *string
	dirPath = flag.String("dirPath", "", "Path to dir with files")
	// строка, которую будем искать в файлах
	substrRequired = "Hello World!"
)

func main() {
	// парсим то, то что нам передали
	flag.Parse()
	// проверяем, что нам передали директорию
	if *dirPath == "" {
		/* если не передали, то бух в панику
		не самая хорошая практика, но мы только для примера
		по хорошему - написать ошибку и выйти
		например так:
			fmt.Println("dirPath shouldn't be empty")
			os.Exit(2)
		*/

		panic("dirPath shouldn't be empty")
	}

	/* тк ниже есть функция getStats, которая возвращает структуру, которая nil
	 то при обращении к полям струкруты получим неявную панику
	обработаем панику
	*/
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic occurred:", err)
		}
	}()

	// вызовем функцию, которая пробежит по директории и вернет структуру со статистикой + ошибку
	stats, err := getStats(*dirPath)
	// обработаем ошибку
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	// выведем стастику
	fmt.Println(stats.Count, stats.FileNames)
}

// описание cтруктуры со статистикой
type HelloWorldStats struct {
	Count     int
	FileNames []string
}

// функция, которая принимает путь до файла, читает его, закрывает и ищет подстроку и возвращает есть ли подстрока в файде
func hasSubstring(filePath string) (bool, error) {
	// открываем файл на чтение
	file, err := os.Open(filePath)
	// закрыть файл, когда закончится выполнение
	defer file.Close()
	if err != nil {
		// не смогли открыть - ошибка с "оберточкой"
		return false, fmt.Errorf("failed to open for substring: %w", err)
	}

	// читаем файл
	b, err := ioutil.ReadAll(file)
	if err != nil {
		return false, err
	}

	// конвертнем считаный файл из байтов в строку
	fileCont := string(b)

	// поищем в файле искомую строку и вернем результат
	return strings.Contains(fileCont, substrRequired), nil
}

// функция, которая решает нашу задачу - получает директорию и создает для директории информацию со статистикой
func getStats(dirPath string) (*HelloWorldStats, error) {
	// создаем экземпляр где будем хранить статистику и инициализируем
	stats := &HelloWorldStats{
		Count:     0,
		FileNames: make([]string, 0),
	}

	// узнаем какие файлы есть в директории
	allFiles, err := ioutil.ReadDir(dirPath)
	// обрабатываем ошибку и возвращаем ошибку нашего типа
	if err != nil {
		return nil, ErrFailedToGetStats{path: dirPath, innerError: err.Error()}
	}
	// бежим по списку файлов из директории
	for _, f := range allFiles {
		filePath := dirPath + "/" + f.Name()
		// запускаем поиск подстроки для конкретного файла
		ok, err := hasSubstring(filePath)
		// не забываем про ошибки
		if err != nil {
			return nil, ErrFailedToGetStats{path: filePath, innerError: err.Error()}
		}
		// если подстрока найдена, добавим инфу в статистику
		if ok {
			stats.Count++
			stats.FileNames = append(stats.FileNames, f.Name())
		}
	}
	// когда обработали все файлы, возвращаем статистику
	return stats, nil
}

// структура для нашей ошибки
type ErrFailedToGetStats struct {
	path       string
	innerError string
}

// реализация метода Error(), что бы реализовать интерфейс error
func (e ErrFailedToGetStats) Error() string {
	return fmt.Sprintf("failed to getStats for path: %s reason: %s", e.path, e.innerError)
}

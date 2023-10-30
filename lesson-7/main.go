package main

import (
	"fmt"
	"reflect"
)

// Написать функцию, которая принимает на вход структуру in (struct или кастомную struct) и
// values map[string]interface{} (key - название поля структуры, которому нужно присвоить
// value этой мапы). Необходимо по значениям из мапы изменить входящую структуру in с
// помощью пакета reflect. Функция может возвращать только ошибку error.

// создаю пустую структуру, которую потом буду менять
var inSt struct{}

func main() {
	// создаю мапу в нужном формате и присваиваю ей значение
	values := make(map[string]interface{})
	values["one"] = 1

	// вызываю функцию, которой передаю указатель на структуру и копию мапы
	structEdit(&inSt, values)
	// вывожу резульатат для проверки, должно быть {one 1}
	fmt.Println(inSt) // {}
}

func structEdit(inStruct *struct{}, valuesMap map[string]interface{}) {
	// присваиваем val значение мапы, напрямую с ней работать нельзя, т.к.
	// у нее тип не reflect.value и соответственно нет методов рефлексии
	val := reflect.ValueOf(valuesMap)

	// присваиваем in значение пустой структуры, которую передали по указателю
	// при этом разыменовываем ее и через Elem берем значения через указатель
	// синактсис конечно странновато выглядит - сахара не хватает
	in := reflect.ValueOf(&inStruct).Elem()

	// здесь я беру цикл, цель которого - это получить значение key
	// ну и посмотреть, а как вообще работать с key, полученной таким образом
	// итого e - это key, а v - это значение мапы по этому key
	for _, e := range val.MapKeys() {
		v := val.MapIndex(e).Elem()

		// тут мы программно создаем структуру, определяем типы данных
		t := reflect.StructOf([]reflect.StructField{
			{
				Name: "Key",
				Type: reflect.TypeOf(e.Interface()), // string
			},
			{
				Name: "Value",
				Type: reflect.TypeOf(v.Interface()), // int (это странно, ведь был interface{} у мапы
			},
		})

		fmt.Println(t) // struct { Key string; Value int }

		// тут мы создаем структуру и передаю значения
		// вместо того чтобы напрямую указывать тип, можно сделать варианты по типам
		in = reflect.New(t).Elem()
		in.Field(0).SetString(e.String())
		in.Field(1).SetInt(v.Int())

		fmt.Println(in) // {one 1}
	}
}

//func main() {
//	v := struct {
//		FieldString string `json:"field_string"`
//		FieldInt    int
//	}{
//		FieldString: "stroka",
//		FieldInt:    107,
//	}
//
//	PrintStruct(v)
//}
//
//func PrintStruct(in interface{}) {
//	if in == nil {
//		return
//	}
//
//	val := reflect.ValueOf(in)
//
//	if val.Kind() == reflect.Ptr {
//		val = val.Elem()
//	}
//
//	if val.Kind() != reflect.Struct {
//		return
//	}
//
//	for i := 0; i < val.NumField(); i++ {
//		typeField := val.Type().Field(i)
//		if typeField.Type.Kind() == reflect.Struct {
//			log.Printf("nested field: %v", typeField.Name)
//			PrintStruct(val.Field(i).Interface())
//			continue
//		}
//		log.Printf("\tname=%s, type=%s, value=%v, tag=`%s`\n",
//			typeField.Name,
//			typeField.Type,
//			val.Field(i),
//			typeField.Tag,
//		)
//	}
//}

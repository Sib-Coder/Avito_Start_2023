package endpoint

import (
	"encoding/json"
	"fmt"
)

// пример идеи как буду работать с slug коррект
type MyData struct {
	MyArray []int `json:"myarray"`
}

func main() {

	jsonData := []byte(`{"myarray":[1, 2, 3, 4, 5]}`)
	var data MyData
	err := json.Unmarshal(jsonData, &data)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}
	fmt.Println(data.MyArray)

	my_mas := data.MyArray
	fmt.Println(my_mas[1])

	for k, el := range my_mas {
		fmt.Println("key:", k, "element:", el)
	}
}

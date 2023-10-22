package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func main() {
	output, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	outputStr := string(output)
	fmt.Println(output)
	arr := strings.Split(outputStr[:len(outputStr)-1], "\n") //[:len(outputStr)-1] это штука убирает последний элемент из массива, Вчера я делал без этого, но не помню как, мне подсказывали(

	x := len(arr[0]) //
	y := len(arr)

	fmt.Println(x, " ", y)
	fmt.Println(outputStr)

	//
	//
	//Код ниже запускает каждый quad. exec - это стандартная библеотека, и запрет на ее использование не было))))

	cmd1 := exec.Command("./QuadA", strconv.Itoa(x), strconv.Itoa(y))
	stdout1, err := cmd1.Output()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(stdout1))

	cmd2 := exec.Command("./QuadB", strconv.Itoa(x), strconv.Itoa(y))
	stdout2, err := cmd2.Output()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(stdout2))

	cmd3 := exec.Command("./QuadC", strconv.Itoa(x), strconv.Itoa(y))
	stdout3, err := cmd3.Output()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(stdout3))

	cmd4 := exec.Command("./QuadD", strconv.Itoa(x), strconv.Itoa(y))
	stdout4, err := cmd4.Output()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(stdout4))

	cmd5 := exec.Command("./QuadE", strconv.Itoa(x), strconv.Itoa(y))
	stdout5, err := cmd5.Output()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(stdout5))

	// Это пример того - как я хочу воспроизвести логику  -> сравнить результат с каждым quad и сделать вывод
	if string(stdout3) == outputStr {
		fmt.Println("!")
	}
}

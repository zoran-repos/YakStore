package main

import (
	"encoding/xml"
	"fmt"
	f "herd/functions"
	"io/ioutil"
	"os"
	"strconv"
	h "herd/models"
)

func main() {

	var days int64
	var herd h.Herd_Xml
	var tillDateStock, milkStock, oldInAges float64
	var skinStock, tillDateSkinStock int
	var herds []string 

	if len(os.Args) != 3 {
		fmt.Println("Usage:", os.Args[0], "NumberOfDays", "fileName")
		return
	}
	xmlFile, err := os.Open(os.Args[2])
	days, _ = strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	defer xmlFile.Close()
	byteValue, _ := ioutil.ReadAll(xmlFile)

	xml.Unmarshal(byteValue, &herd)

	for i := 0; i < len(herd.Labyak); i++ {
		oldInAges, _ = strconv.ParseFloat(herd.Labyak[i].Age, 64)
		ages, _ := strconv.ParseFloat(herd.Labyak[i].Age,64) 
		ages = (ages * 100.0 + float64(days))/100.0
		herds=  append(herds, herd.Labyak[i].Name + " " + fmt.Sprintf("%v", ages) +  " years old")

		tillDateStock = f.GetTotalMilk(int(days), float64(oldInAges))
		milkStock += tillDateStock

		tillDateSkinStock = f.GetSkin(int(days), oldInAges)
		skinStock += tillDateSkinStock
	}
	fmt.Printf("Output for T = %d \n\n", days)
	fmt.Println("In stock:")
	fmt.Printf("    %.3f \n", milkStock)
	fmt.Printf("    %d \n", skinStock)
	fmt.Println("Herd:")
	for _, herdLine := range herds {
        fmt.Println("    " + herdLine)
	}

}

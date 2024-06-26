package main

import (
	"fmt"
	cli "main/pkg/cli"
	arr "main/pkg/plot/arr"
	draw "main/pkg/plot/draw"
)

func main() {

	directory := "."
	patternFileName := "standard.txt"
	//patternFileName := "thinkertoy.txt"
	chLength := 8
	var firstCh byte = ' '
	var lastCh byte = '~'
	//----------------------------------------
	patternContent, errPattern := cli.ReadFileToStr(directory, patternFileName)

	if errPattern != nil {
		fmt.Println("Error : loading pattern file failed")
		return
	}
	//--------------------------------
	patternMap, errPatternMap := arr.ContentToMap(patternContent, chLength, firstCh, lastCh)

	if errPatternMap != nil {
		fmt.Println(errPatternMap)
		return
	}
	//--------------------------------------------
	cliStr, errClistr := cli.ReadCli()

	if errClistr != nil {
		//fmt.Println(errClistr)
		return
	}
	//----------------------------------------------

	finalStr := draw.StrMaker(cliStr, patternMap, chLength)
	fmt.Println(finalStr)
	//-----------------------------------------
}

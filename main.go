package main 

import (
	"encoding/json"
	"fmt"
	"os"
	"io/ioutil"
	"math"
	"strconv"
)

type Rectangle struct {
	X int `json:"x"`
	Y int `json:"y"`
	W int `json:"w"`
	H int `json:"h"`
}


type Rects struct {
	Rects []Rectangle `json:"rects"`
}

type RectangeDimension struct {
	X int
	Y int
	W int
	H int
}

type Data struct {
	RecNumber []int
	RecDimension RectangeDimension 
}

func checkIntersectioin(r1tx1 int, r1ty1 int, r1bx2 int, r1by2 int, r2tx1 int, r2ty1 int, r2bx2 int, r2by2 int) (status bool, X1 float64, Y1 float64, W float64, H float64) {
	// fmt.Println("Check intersection input", r1tx1, r1ty1, r1bx2, r1by2, r2tx1, r2ty1, r2bx2, r2by2)
	// First, Check two rectangles can intersect or not.
	if r2ty1 > r1by2 {
		// fmt.Println("Here intersect cannot possible.")
		return false, 0, 0, 0, 0
	} else if r2tx1 > r1bx2 {
		return false, 0, 0, 0, 0
	} else {
		// fmt.Println("Here, possible")	
	
		maxX := math.Max(float64(r1tx1), float64(r2tx1))
		maxY := math.Max(float64(r1ty1), float64(r2ty1))

		minX := math.Min(float64(r1bx2) , float64(r2bx2))
		minY := math.Min(float64(r1by2), float64(r2by2))
		
		// fmt.Println(maxX, maxY, math.Abs(maxX - minX), math.Abs(maxY - minY))
	
		return true, maxX, maxY, math.Abs(maxX - minX), math.Abs(maxY - minY)
	}
}

func NewIntersect(rect []Rectangle) (emptyJson bool, arrayOfData []Data) {

	arrayOfData = []Data{}
	secondArrayOfIntermediateData := []Data{}

	r1tx1 := 0
	r1ty1 := 0
	r1bx2 := 0
	r1by2 := 0

	r2tx1 := 0
	r2ty1 := 0
	r2bx2 := 0
	r2by2 := 0

	itx1 := 0.0
	ity1 := 0.0

	iw := 0.0
	ih := 0.0
	status := false

	emptyJson = false

	// Iterate over full rectangle array.
	for key, _ := range rect {

		// fmt.Println(key)

		// Iterate from next index to end of array.
		for innerKey := key +1; innerKey< len(rect) ; innerKey++ {

			r1tx1 = rect[key].X
			r1ty1 = rect[key].Y
			r1bx2 = rect[key].X + rect[key].W
			r1by2 = rect[key].Y + rect[key].H

			r2tx1 = rect[innerKey].X
			r2ty1 = rect[innerKey].Y
			r2bx2 = rect[innerKey].X + rect[innerKey].W
			r2by2 = rect[innerKey].Y + rect[innerKey].H

			status, itx1, ity1, iw, ih = checkIntersectioin(r1tx1, r1ty1, r1bx2, r1by2, r2tx1, r2ty1, r2bx2, r2by2)
			if status {
				// fmt.Println("Intersect possible.", itx1, ity1, iw, ih)
				
				var result Data

				result.RecNumber = []int{key, innerKey}
				result.RecDimension.X = int(itx1)
				result.RecDimension.Y = int(ity1)
				result.RecDimension.W = int(iw)
				result.RecDimension.H = int(ih)

				arrayOfData = append(arrayOfData, result)
				

			}

		}

	}


	// fmt.Println("-----")

	for _, value := range arrayOfData {
		// fmt.Println(value.RecNumber, value.RecDimension.X, value.RecDimension.Y, value.RecDimension.W, value.RecDimension.H)

		// Start - Calculate max rectanble number.
		maxRecNumber := 0
		for _, number := range value.RecNumber {
			// fmt.Println(number)
			if number > maxRecNumber {
				maxRecNumber = number
			}
		}
		// End - Calculate max rectanble number.

		// Start - Iteration for multiple rectangles.
		for _, _ = range rect {
			if maxRecNumber == len(rect) - 1 {
				// fmt.Println("Rectanble array len matched.")
				break
			}

			innerKey := maxRecNumber + 1

			r1tx1 = value.RecDimension.X
			r1ty1 = value.RecDimension.Y
			r1bx2 = value.RecDimension.X + value.RecDimension.W
			r1by2 = value.RecDimension.Y + value.RecDimension.H

			r2tx1 = rect[innerKey].X
			r2ty1 = rect[innerKey].Y
			r2bx2 = rect[innerKey].X + rect[innerKey].W
			r2by2 = rect[innerKey].Y + rect[innerKey].H


			for innerKey := maxRecNumber +1; innerKey< len(rect) ; innerKey++ {

				
				status, itx1, ity1, iw, ih = checkIntersectioin(r1tx1, r1ty1, r1bx2, r1by2, r2tx1, r2ty1, r2bx2, r2by2)
				if status {
					// fmt.Println("intersect between ")
					// fmt.Println(value.RecNumber, innerKey)
					r2tx1 = int(itx1)
					r2ty1 = int(ity1)
					r2bx2 = int(itx1) + int(iw)
					r2by2 = int(ity1) + int(ih)
					// fmt.Println("-----")

					value.RecNumber = append(value.RecNumber, innerKey)

					var result Data

					result.RecNumber = value.RecNumber
					result.RecDimension.X = int(itx1)
					result.RecDimension.Y = int(ity1)
					result.RecDimension.W = int(iw)
					result.RecDimension.H = int(ih)

					secondArrayOfIntermediateData = append(secondArrayOfIntermediateData, result)
				}
			}
			if innerKey == len(rect) - 1 {
				// fmt.Println("Inner key = len matched.")
				break
			}

		} 

		// End - Iteration for multiple rectangles

	}

	if len(rect) == 0 {
		// fmt.Println("Not readable JSON file.")
		emptyJson = true
	} else {
		arrayOfData = append(arrayOfData, secondArrayOfIntermediateData...)
	}
	
	return emptyJson, arrayOfData

}


func main() {
	args := os.Args
	
	if len(args) > 1 {
		fileData, err := ioutil.ReadFile(args[1])
		if err != nil {
			fmt.Println("Error while reading JSON file.")
		}

		var rects Rects
		json.Unmarshal(fileData, &rects)
		// fmt.Println("rects", rects, len(rects.Rects))
		tmpRects := rects.Rects
		if len(rects.Rects) > 10 {
			tmpRects = rects.Rects[:10]
		} 

		status, result := NewIntersect(tmpRects)
		// fmt.Println("result", status,  result)
		if status {
			fmt.Println("JSON file not in proper format.")
		} else {
			fmt.Println("Input:")
			for key, value := range tmpRects {
				// fmt.Println(key, value)
				fmt.Printf("\t %d: Rectangle at (%d,%d), w=%d, h=%d. \n", key + 1, value.X, value.Y, value.W, value.H)
			}

			fmt.Println("Intersections:")
			// arrayOfData = append(arrayOfData, secondArrayOfIntermediateData...)
			for key, value := range result {

				recSequence := ""
				for recKey, recNumber := range value.RecNumber {
					if recKey == len(value.RecNumber) - 2 {
						recSequence += strconv.Itoa(recNumber + 1) + " and "
					} else {
						recSequence += strconv.Itoa(recNumber + 1) +  ", "
					} 
					// else {
					// 	recSequence += strconv.Itoa(recNumber + 1) + " at"
					// }
				}
				fmt.Printf("\t %d Between rectangle %sat (%d, %d), w=%d, h=%d. \n", key + 1, recSequence, value.RecDimension.X, value.RecDimension.Y, value.RecDimension.W, value.RecDimension.H)
				// fmt.Println(key, value)
			}
		}
	} else {
		fmt.Println("File name not found in arguments.")
	}
	
}
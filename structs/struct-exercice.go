package main

/* create a program to calculate the area and the perimeter of a rectangle */

import (
	"fmt"
);

type Rectangle struct{
	Length int;
	Width int;
}


func calculateRectangleArea(RectangleCopy Rectangle) int{
	return RectangleCopy.Width * RectangleCopy.Length
}

func calculateRectanglePerimeter(StructCopy Rectangle)int{
	return (StructCopy.Length + StructCopy.Width )* 2
}

func main(){
   rectangle := Rectangle{20, 40};
   fmt.Println("the perimeter of the rectangle is", calculateRectanglePerimeter(rectangle),"cm");
   fmt.Println("the area of the rectangle is", calculateRectangleArea(rectangle), "cm");

   //double the size of the rectangle;
   rectangle.Length *= 2;
   rectangle.Width *= 2;
   fmt.Println("the new perimeter of the rectangle is", calculateRectanglePerimeter(rectangle),"cm");
   fmt.Println("the new area of the rectangle is", calculateRectangleArea(rectangle), "cm");


}
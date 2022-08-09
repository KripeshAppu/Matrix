package main
import "fmt"
func main(){
	var  row, col, i, j, temp, value int
	var a[50][50] int
	fmt.Print("Enter number of rows:")
	fmt.Scan(&row)
	fmt.Print("Enter number of columns:")
	fmt.Scan(&col)
	if row == col {
				a[0][0]=1
				a[row-1][col-1]=row*col
				value=2
				for i=1; i<row; i++{
					temp=i
					for j=0; j<=i; j++{
						a[temp][j]=value
						temp--
						value++
					}
				}
				for j=1; j<col-1; j++{
					temp=j
					for i=row-1; i>=j; i--{
						a[i][temp]=value
						temp++
						value++
					}
				}
	}else if row<col{
				a[0][0]=1
				a[row-1][col-1]=row*col
				value=2
				for i=1; i<row; i++{
					temp=i
					for j=0; j<=i; j++{
						a[temp][j]=value
						temp--
						value++
					}
				}
				for j=1; j<col-1; j++{
					temp=j
					for i=row-1; i>=j-1; i--{
						a[i][temp]=value
						temp++
						value++
					} 
				}
	}else{
		a[0][0]=1
		a[row-1][col-1]=row*col
		value=2
		for i=1; i<row-1; i++{
			temp=i
			for j=0; j<=i; j++{
				a[temp][j]=value
				temp--
				value++
			}
		}
		for j=0; j<col-1; j++{
			temp=j
			for i=row-1; i>=j+1; i--{
				a[i][temp]=value
				temp++
				value++
			} 
		}
		} 
	fmt.Println("Matrix:")
	for i=0; i<row; i++{
		for j=0; j<col; j++{
			fmt.Print(a[i][j],"\t")
		}
		fmt.Print("\n")
	}
}
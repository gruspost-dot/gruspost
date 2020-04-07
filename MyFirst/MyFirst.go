package MyFirst

import "fmt" 

// Hi returns a friendly greeting
func Hellos(name string) string {
   return fmt.Sprintf("Hi, %s", name)
}
package main

import(
	"fmt"
	"sort"
)

func longestCommonPrefix(strs []string) string {
    if len(strs) == 0 {
        return ""
    }
    sort.Strings(strs)
    first, last := strs[0], strs[len(strs)-1]
    result := ""
    for i := 0; i < len(first); i++ {
        if i < len(last) && first[i] == last[i] {
            result += string(first[i])
        } else {
            break
        }
    }
    return result
}

func main(){
	array := []string{"flow", "flower", "flight"}
	sort.Strings(array)
	fmt.Println(array)
	first, last := array[0], array[len(array)-1]
	fmt.Println(first, last, len(last))
	fmt.Println(longestCommonPrefix(array))
}
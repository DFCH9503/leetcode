//using native js methods and functions

function isPalindrome(x){

    res = parseInt(x.toString().split("").reverse().join(""))

    return res == x
}

//using other approach

// function isPalindrome(x){

    
// }

console.log("answer to the first example is:", isPalindrome(121))

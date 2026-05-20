package main

import (

"demo/3-struct/bins"

)




func main() {
bin1, _ := bins.InitBin("123", true, "penis")
bins.InitBins([]bins.Bin{*bin1})
}


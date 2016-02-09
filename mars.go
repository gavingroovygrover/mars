// Copyright 2016 Gavin "Groovy" Grover. All rights reserved.
// Use of this source code is governed by the same BSD-style
// license as Go that can be found in the LICENSE file.

package main

import(
  "fmt"
)

type MonthData []struct{ ds int; nm string }

var earthMths = MonthData {
   {31, "Jan"},  {62, "Feb"},  {90, "Mar"}, {121, "Apr"}, {151, "May"}, {182, "Jun"},
  {212, "Jul"}, {243, "Aug"}, {274, "Sep"}, {304, "Oct"}, {335, "Nov"}, {365, "Dec"},
} 

var marsMths = MonthData {
   {61, "Ari"}, {127, "Tau"}, {193, "Gem"}, {258, "Can"}, {318, "Leo"}, {372, "Vir"},
  {422, "Lib"}, {469, "Sco"}, {515, "Sag"}, {562, "Cap"}, {613, "Aqu"}, {669, "Pis"},
}

const(
  day = 24 * 60 * 60
  sol = day + 39 * 60 + 35.24409
  daysPerSol = sol/day // approx 1.0274912510 Earth days/sol 
  marsYear = 668.5991
  earthYear = 365.2564
  numYrs = 25
)

func main(){
  var mthToggle bool
  var monthsUsed = map[string]int{}
  type MthStack struct{ arr []string; top int }
  var mthStack MthStack = MthStack{[]string{"", "", "", "", ""}, 5}
  var ma, mm, ms int = 1, 1, 0 //Martian annum, month, sol
  var esta, edtm, edty float64 = 0.0, 0.0, 0.0 //elapsed sols this annum, earth-days this month, earth-days this year

  fmt.Print("月: ")
  for i:= 1; i <= 19; i++ { fmt.Printf("%4d  ", i) }
  fmt.Printf("\n%2d年:", ma)
  for ; ma <= numYrs; ms++ {
    if edtm > float64(ms) + 1 && ms % 7 == 6 {
      var mthNm string
      if mthToggle {
        mthNm= "Pis"
        mlab: for i:= 0; i < 12; i++ {
          if esta <= float64(marsMths[i].ds) {
            mthNm= marsMths[(i+11)%12].nm
            break mlab
          }
        }
      } else {
        mthNm= "Dec"
        elab: for i:= 0; i < 12; i++ {
          if edty <= float64(earthMths[i].ds) {
            mthNm= earthMths[(i+11)%12].nm
            if mthNm == mthStack.arr[mthStack.top - 5] {
              mthNm= earthMths[i%12].nm
            }
            break elab
          }
        }
        if mthNm == "Dec" && mthNm == mthStack.arr[mthStack.top - 5] {
          mthNm= "Jan"
        }
        mthStack.arr = append(mthStack.arr, mthNm)
        mthStack.top++
      }
      fmt.Printf(" %s", mthNm)
      monthsUsed[mthNm]++
      mthToggle = !mthToggle
      if ms == 34 { fmt.Printf("  ") } else { fmt.Printf("+ ") }
      if esta > marsYear {
        ma++; mm = 0
        esta -= marsYear
        fmt.Printf("\n%2d年:", ma)
      }
      edtm -= float64(ms) + 1
      if edty > earthYear {
        edty -= earthYear
      }
      mm++; ms = 0
    }
    edtm += daysPerSol
    edty += daysPerSol
    esta++
  }
  fmt.Println()
  for k, v:= range monthsUsed { fmt.Printf("%s:%2d, ", k, v) }
  fmt.Println()
}


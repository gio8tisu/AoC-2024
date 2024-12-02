package utils

func AbsDiffUint(x, y uint64) uint64 {
   if x < y {
      return y - x
   }
   return x - y
}

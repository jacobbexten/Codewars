package kata

func SquaresInRect(lng int, wdth int) []int {
  longestSide := 0
  squares := []int{}
  
  if lng == wdth {
    return nil
  }
  
  for lng > 0 && wdth > 0 {
    if lng < wdth {
      longestSide = lng
      wdth -= lng
    } else {
      longestSide = wdth
      lng -= wdth
    }

    squares = append(squares, longestSide)
    
  }
  
return squares
  
}

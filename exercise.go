package exercise

import "bytes"

func Play(input []string) []string{

  p1 := MakeShotMap(input[0])
  p2 := MakeShotMap(input[1])

  p1Score := CalcPlayerScore(p1,p2)
  p2Score := CalcPlayerScore(p2,p1)

  p1Letters := CalcLeftOverLetters(p1, p2)
  p2Letters := CalcLeftOverLetters(p2, p1)

  return []string{p1Letters,p2Letters,WhoWon(p1Score,p2Score)}
}

func CalcLeftOverLetters(player map[string]int, oponent map[string]int) string {
  var buffer bytes.Buffer

  for k, playerValue := range player {
    if opValue, ok := oponent[k]; ok {
      numberOfLetters := (playerValue - opValue)
      if numberOfLetters > 0 {
        for i := 0; i < numberOfLetters; i ++ {
          buffer.WriteString(k)
        }
      }
    }else {
      for i := 0; i < playerValue; i ++ {
        buffer.WriteString(k)
      }
    }
  }
  return buffer.String()
}

func WhoWon(player1Score int, player2Score int) string{
  scoreDiff := player1Score - player2Score
  if scoreDiff == 0 {
    return "tie"
  }
  if scoreDiff > 0 {
    return "p1"
  }
  return "p2"
}

func MakeShotMap(input string) map[string]int {
  result := make(map[string]int)
  for _, c := range input {
    if _, ok := result[string(c)]; ok {
      result[string(c)]++
    }else {
      result[string(c)] = 1
    }
  }
  return result
}

func CalcPlayerScore(player map[string]int, oponent map[string]int) int {
  result := 0

  for k, playerValue := range player {
    if opValue, ok := oponent[k]; ok {
      if (playerValue > opValue){
        result = result + (playerValue - opValue)
      }
    }else {
      result = result + playerValue
    }
  }
  return result
}

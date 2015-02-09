package exercise

import (
  . "github.com/smartystreets/goconvey/convey"
  "testing"
)

func TestPlay(t *testing.T){

  Convey("Test tie game short words", t, func(){
    input := []string{"hat", "cat"}
    result := Play(input)

    Convey("hat v cat results in tie", func(){
      So(result[2], ShouldEqual, "tie")
    })

    Convey("hat v cat left and right values", func(){
      So(result[0], ShouldEqual,"h")
      So(result[1], ShouldEqual,"c")
    })

  })

  Convey("Test tie game long words", t, func(){
    input := []string{"million", "billion"}
    result := Play(input)

    Convey("hat v cat results in tie", func(){
      So(result[2], ShouldEqual, "tie")
    })

    Convey("hat v cat left and right values", func(){
      So(result[0], ShouldEqual,"m")
      So(result[1], ShouldEqual,"b")
    })
  })

  Convey("Given p1 'longbow' p2 'blowup'", t, func(){

    input := []string{"longbow", "blowup"} //ngo up

    Convey("When the game is Played", func(){
      result := Play(input)
      Convey("Player 1 wins", func() {
        So(result[2], ShouldEqual, "p1")
      })

      Convey("the right left over letters are returned", func(){
        So(result[0], ShouldContainSubstring,"n")
        So(result[0], ShouldContainSubstring,"g")
        So(result[0], ShouldContainSubstring,"o")
        So(len(result[0]), ShouldEqual,3)

        So(result[1], ShouldContainSubstring,"u")
        So(result[1], ShouldContainSubstring,"p")
        So(len(result[1]), ShouldEqual,2)
      })
    })

  })

}

func TestLeftOverLetters(t *testing.T) {
  Convey("Given p1 'foo' p2 'bar' maps", t, func(){
    p1 := map[string]int{
    "f": 1,
    "o": 2,
    }

    p2 := map[string]int{
    "b": 1,
    "a": 1,
    "r": 1,
    }

    Convey("it returns foo for player 1", func(){
      result := CalcLeftOverLetters(p1, p2)
      So(result, ShouldContainSubstring, "f")
      So(result, ShouldContainSubstring, "oo")
    })
    Convey("it returns bar for player 2", func(){
      result := CalcLeftOverLetters(p2, p1)
      So(result, ShouldContainSubstring, "b")
      So(result, ShouldContainSubstring, "a")
      So(result, ShouldContainSubstring, "r")
    })

  })

  Convey("Given p1 'hat' p2 'bar' maps", t, func(){
    p1 := map[string]int{
    "h": 1,
    "a": 1,
    "t": 1,
    }

    p2 := map[string]int{
    "b": 1,
    "a": 1,
    "r": 1,
    }

    Convey("it returns 'ha' for player 1", func(){
      result := CalcLeftOverLetters(p1, p2)
      So(result, ShouldContainSubstring, "h")
      So(result, ShouldContainSubstring, "t")
    })
    Convey("it returns ba for player 2", func(){
      result := CalcLeftOverLetters(p2, p1)
      So(result, ShouldContainSubstring, "b")
      So(result, ShouldContainSubstring, "r")
    })

  })

}

func TestWhoWon(t *testing.T){

  Convey("Given a that player 1 score higher than player 2", t, func(){
    p1 := 3
    p2 := 2

    Convey("it returns player 1 win", func(){
      result := WhoWon(p1, p2)
      So(result, ShouldEqual, "p1")

      })
    })

    Convey("Given a that player 2 score higher than player 1", t, func(){
      p1 := 4
      p2 := 5

      Convey("it returns player 2 win", func(){
        result := WhoWon(p1, p2)
        So(result, ShouldEqual, "p2")

      })
    })

      Convey("Given a that player 1 score equal to player 2", t, func(){
        p1 := 4
        p2 := 4

        Convey("it returns tie", func(){
          result := WhoWon(p1, p2)
          So(result, ShouldEqual, "tie")

        })
      })
}


func TestMakeShotMap(t *testing.T){

  Convey("Given a string with no repeating letters", t, func(){
    input := "bar"

    Convey("it returns the correct result", func(){
      exp := make(map[string]int)
      exp["b"] = 1
      exp["a"] = 1
      exp["r"] = 1
      result := MakeShotMap(input)
      So(len(result), ShouldEqual, 3)

      for k, v := range exp {
        So(result[k], ShouldEqual, v)
      }
    })
  })

  Convey("Given a string 'longbow'", t, func(){
    input := "longbow"

    Convey("it returns the correct result", func(){
      exp := make(map[string]int)
      exp["l"] = 1
      exp["o"] = 2
      exp["n"] = 1
      exp["g"] = 1
      exp["b"] = 1
      exp["w"] = 1
      result := MakeShotMap(input)
      So(len(result), ShouldEqual, 6)

      for k, v := range exp {
        So(result[k], ShouldEqual, v)
      }
    })
  })

  Convey("Given a string 'blowup'", t, func(){
    input := "blowup"

    Convey("it returns the correct result", func(){
      exp := map[string]int{
      "b": 1,
      "l": 1,
      "o": 1,
      "w": 1,
      "u": 1,
      "p": 1,
      }
      result := MakeShotMap(input)
      So(len(result), ShouldEqual, 6)

      for k, v := range exp {
        So(result[k], ShouldEqual, v)
      }
    })
  })

  Convey("Given a string with repeating letters", t, func(){
    input := "fooo"

    Convey("it returns the correct result", func(){
      exp := make(map[string]int)
      exp["f"] = 1
      exp["o"] = 3
      result := MakeShotMap(input)
      So(len(result), ShouldEqual, 2)

      for k, v := range exp {
        So(result[k], ShouldEqual, v)
      }
    })
  })
}

func TestCalcPlayerScore(t *testing.T){
  Convey("Given p1 'longbow' p2 'blowup' maps", t, func(){
    p1 := map[string]int{
    "l": 1,
    "o": 2,
    "n": 1,
    "g": 1,
    "b": 1,
    "w": 1,
    }

    p2 := map[string]int{
    "b": 1,
    "l": 1,
    "o": 1,
    "w": 1,
    "u": 1,
    "p": 1,
    }

    Convey("it returns the correct p1 score", func(){
      result := CalcPlayerScore(p1, p2)
      So(result, ShouldEqual, 3)
    })
    Convey("it returns the correct p2 score", func(){
      result := CalcPlayerScore(p2, p1)
      So(result, ShouldEqual, 2)
    })

  })

  Convey("Given p1 'foo' p2 'bar' maps", t, func(){
    p1 := map[string]int{
    "f": 1,
    "o": 2,
    }

    p2 := map[string]int{
    "b": 1,
    "a": 1,
    "r": 1,
    }

    Convey("it returns the correct p1 score", func(){
      result := CalcPlayerScore(p1, p2)
      So(result, ShouldEqual, 3)
    })
    Convey("it returns the correct p2 score", func(){
      result := CalcPlayerScore(p2, p1)
      So(result, ShouldEqual, 3)
    })

  })

  Convey("Given p1 'joe' p2 'joe' maps", t, func(){
    p1 := map[string]int{
    "j": 1,
    "o": 1,
    "e": 1,
    }
    p2 := map[string]int{
    "j": 1,
    "o": 1,
    "e": 1,
    }

    Convey("it returns the correct p1 score", func(){
      result := CalcPlayerScore(p1, p2)
      So(result, ShouldEqual, 0)
    })
    Convey("it returns the correct p2 score", func(){
      result := CalcPlayerScore(p2, p1)
      So(result, ShouldEqual, 0)
    })

  })

  Convey("Given p1 'looser' p2 'winning' maps", t, func(){
    p1 := map[string]int{
    "l": 1,
    "o": 2,
    "s": 1,
    "e": 1,
    "r": 1,
    }
    p2 := map[string]int{
    "w": 1,
    "i": 2,
    "n": 3,
    "g": 1,
    }

    Convey("it returns the correct p1 score", func(){
      result := CalcPlayerScore(p1, p2)
      So(result, ShouldEqual, 6)
    })
    Convey("it returns the correct p2 score", func(){
      result := CalcPlayerScore(p2, p1)
      So(result, ShouldEqual, 7)
    })

  })
}

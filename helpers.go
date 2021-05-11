package main

type Move struct {
	XDelta int
	YDelta int
	Text   string
}

var (
	Noop  = Move{}
	Up    = Move{XDelta: 0, YDelta: 1, Text: "up"}
	Down  = Move{XDelta: 0, YDelta: -1, Text: "down"}
	Left  = Move{XDelta: -1, YDelta: 0, Text: "left"}
	Right = Move{XDelta: 1, YDelta: 0, Text: "right"}
)

func Dist(a, b Coord) int {
	return Abs(a.X-b.X) + Abs(a.Y-b.Y)
}

func Abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

type Result string

const (
	Death      Result = "😵"
	MaybeDeath Result = "🤷‍♂️"
	Fed        Result = "🍔"
	Alive      Result = "🥳"
)

// Valid returns true if a move will not result in immediate death, false
// otherwise.
func MoveResult(g GameRequest, m Move) Result {
	at := g.You.Head
	going := Coord{at.X + m.XDelta, at.Y + m.YDelta}

	// 1. See if we are going off the board
	if going.X >= g.Board.Width || going.X < 0 || going.Y >= g.Board.Height || going.Y < 0 {
		return Death
	}

	// 2. Check if my own body is here
	for _, coord := range g.You.Body {
		// TODO: Eventually account for moving where my tail was when length is longer than 2. The naive approach gets weird when our size is 2 and we go up, then immediately try to go down where our tail used to be, so it needs to be a hair more sophisticated 😫.
		if going == coord {
			return Death
		}
	}

	// 3. See if any snakes could go to this spot. If so, MaybeDeath.
	for _, s := range g.Board.Snakes {
		if s.Health == 0 {
			continue
		}
		if s.ID == g.You.ID {
			continue
		}
		if Dist(going, s.Head) == 1 {
			return MaybeDeath
		}
	}

	// 4. See if a snake or tail is there. If a tail, we need to see if that snake
	// could eat.
	for _, s := range g.Board.Snakes {
		if s.ID == g.You.ID {
			continue
		}
		for i, coord := range s.Body {
			if going == coord {
				if i == len(s.Body)-1 {
					if CanEat(s.Head, g.Board.Food) {
						return MaybeDeath
					}
					return Alive
				}
				return Death
			}
		}
	}

	// 5. No snakes are going here and no bodies are here. See if food is here.
	for _, f := range g.Board.Food {
		if going == f {
			return Fed
		}
	}

	// 6. Think that's it? Not sure.
	return Alive
}

// CanEat returns true if a piece of food is within 1 move of the provided
// head location.
func CanEat(head Coord, food []Coord) bool {
	for _, f := range food {
		if Dist(head, f) == 1 {
			return true
		}
	}
	return false
}

// Q: Is health 0 if snake is dead, or does it stop showing up?

// ParseASCII converts something like:
//
//  ---aA--
//  ---a---
//  -*-----
//  -------
//  ---bb--
//  ---b--*
//  ---bB--
//
// Into a GameRequest. This is used for simpler testing, and is definitely
// not exhaustive. Eg some ASCII boards are ambiguous and I can't promise
// how those will be parsed.
// func ParseASCII(lines []string) GameRequest {

// }

// TODO: Trap detection
// TODO: Seek food

package types

type GameConf struct {
	Dimensions Vec2
	CanLoop    bool
}

type GameEvent uint8

const (
	GameContinue GameEvent = iota
	GameLost
	AteApple
)

type Direction uint8

const (
	Up Direction = iota
	Down
	Left
	Right
	NoDirection
)

func (d Direction) ToVec2() Vec2 {
	switch d {
	case Up:
		return NewVec2(0, -1)
	case Down:
		return NewVec2(0, 1)
	case Left:
		return NewVec2(-1, 0)
	case Right:
		return NewVec2(1, 0)
	}
	return NewVec2(0, 0)
}

func (d Direction) GetOpposite() Direction {
	switch d {
	case Up:
		return Down
	case Down:
		return Up
	case Left:
		return Right
	case Right:
		return Left
	}
	return NoDirection
}

type Vec2 struct {
	X int
	Y int
}

func NewVec2(x int, y int) Vec2 {
	return Vec2{x, y}
}

func (v Vec2) Add(other Vec2) Vec2 {
	temp := v
	temp.X += other.X
	temp.Y += other.Y
	return temp
}

func (v Vec2) Loop(maxVec Vec2) Vec2 {
	temp := v
	temp.X = (maxVec.X + temp.X) % maxVec.X
	temp.Y = (maxVec.Y + temp.Y) % maxVec.Y
	return temp
}

func (v Vec2) IsOutsideOf(other Vec2) bool {
	return (v.X > other.X || v.Y > other.Y) || (v.X < 0 || v.Y < 0)
}

func (v Vec2) IsEqual(other Vec2) bool {
	return v.X == other.X && v.Y == other.Y
}

type Snake struct {
	Segments  []Vec2
	backlog   uint8
	Direction Direction
}

func NewSnake(startPos Vec2, segments uint8, startDir Direction) Snake {
	segs := make([]Vec2, 1)
	segs[0] = startPos
	snake := Snake{Segments: segs, backlog: segments - 1, Direction: startDir}
	return snake
}

func (s *Snake) Extend() {
	s.Segments = append(s.Segments, s.Segments[len(s.Segments)-1])
}

func (s *Snake) Move(conf GameConf) bool {
	snakeCopy := s

	newHead := snakeCopy.Segments[0].Add(snakeCopy.Direction.ToVec2())
	if conf.CanLoop {
		newHead = newHead.Loop(conf.Dimensions)
	} else if newHead.IsOutsideOf(conf.Dimensions) {
		return false
	}
	snakeCopy.Segments = append([]Vec2{newHead}, s.Segments...)

	if snakeCopy.backlog > 0 {
		snakeCopy.backlog -= 1
	} else {
		snakeCopy.Segments = s.Segments[:len(s.Segments)-1]
	}

	if snakeCopy.IsCollidingSelf() {
		return false
	}
	s = snakeCopy
	return true
}

func (s Snake) IsCollidingSelf() bool {
	for i := 1; i < len(s.Segments); i++ {
		if s.Segments[i].IsEqual(s.GetHead()) {
			return true
		}
	}
	return false
}

func (s Snake) IsOnBody(pos Vec2) bool {
	for i := 1; i < len(s.Segments); i++ {
		if s.Segments[i].IsEqual(pos) {
			return true
		}
	}
	return false
}

func (s Snake) GetHead() Vec2 {
	return s.Segments[0]
}

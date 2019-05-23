// code generated by hasgo. [DO NOT EDIT!]
package player

// =============== Reverse.go =================

// Returns the reversed slice
func (s Players) Reverse() (out Players) {
	for i := len(s) - 1; i >= 0; i-- {
		out = append(out, s[i])
	}
	return
}

// =============== Init.go =================

// Take n-1 elements from a slice, where n = len(list)
func (s Players) Init() (out Players) {
	slicecopy := append([]Player(nil), s...)
	return slicecopy[:len(s)-1]
}

// =============== Last.go =================

// Returns the last element in the slice
// If no element is found, returns the zero-value of the type
func (s Players) Last() (out Player) {
	if len(s) > 0 {
		out = s[len(s)-1]
	}
	return
}

// =============== Tail.go =================

// Take [1 -> n] elements from a slice, where n = len(list)
// Returns an empty slice if there are less than 2 elements in slice
func (s Players) Tail() (out Players) {
	if len(s) <= 1 {
		return
	}
	slicecopy := append([]Player(nil), s...)
	return slicecopy[1:]
}

// =============== Uncons.go =================

func (s Players) Uncons() (head Player, tail Players) {
	return s.Head(), s.Tail()
}

// =============== Filter.go =================

func (s Players) Filter(f func(Player) bool) (out Players) {
	for _, v := range s {
		if f(v) {
			out = append(out, v)
		}
	}
	return
}

// =============== Head.go =================

// Returns the first element in the slice
// If no element is found, returns the zero-value of the type
func (s Players) Head() (out Player) {
	if len(s) > 0 {
		out = s[0]
	}
	return
}

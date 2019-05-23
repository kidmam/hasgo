// code generated by hasgo. [DO NOT EDIT!]
package types

// =============== Init.go =================

// Take n-1 elements from a slice, where n = len(list)
func (s Strings) Init() (out Strings) {
	slicecopy := append([]string(nil), s...)
	return slicecopy[:len(s)-1]
}

// =============== Last.go =================

// Returns the last element in the slice
// If no element is found, returns the zero-value of the type
func (s Strings) Last() (out string) {
	if len(s) > 0 {
		out = s[len(s)-1]
	}
	return
}

// =============== Reverse.go =================

// Returns the reversed slice
func (s Strings) Reverse() (out Strings) {
	for i := len(s) - 1; i >= 0; i-- {
		out = append(out, s[i])
	}
	return
}

// =============== Tail.go =================

// Take [1 -> n] elements from a slice, where n = len(list)
// Returns an empty slice if there are less than 2 elements in slice
func (s Strings) Tail() (out Strings) {
	if len(s) <= 1 {
		return
	}
	slicecopy := append([]string(nil), s...)
	return slicecopy[1:]
}

// =============== Uncons.go =================

func (s Strings) Uncons() (head string, tail Strings) {
	return s.Head(), s.Tail()
}

// =============== Filter.go =================

func (s Strings) Filter(f func(string) bool) (out Strings) {
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
func (s Strings) Head() (out string) {
	if len(s) > 0 {
		out = s[0]
	}
	return
}

// =============== Sum.go =================

func (s Strings) Sum() string {
	var sum string
	for _, v := range s {
		sum += v
	}
	return sum
}

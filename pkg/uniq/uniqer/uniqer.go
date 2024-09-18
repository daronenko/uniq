package uniqer

type Uniqer interface {
	ShouldSave(line string) bool // lazy evaluation
	Finish() bool
}

type DefaultUniqer struct {
	previousLine    string
	hasPreviousLine bool
}

func NewDefaultUniqer() *DefaultUniqer {
	return &DefaultUniqer{}
}

func (u *DefaultUniqer) ShouldSave(line string) bool {
	if !u.hasPreviousLine {
		u.hasPreviousLine = true
		u.previousLine = line
		return false
	}

	if line != u.previousLine {
		u.previousLine = line
		return true
	}

	return false
}

func (u *DefaultUniqer) Finish() bool {
	return true
}

type UniqueUniqer struct {
	previousLine    string
	hasPreviousLine bool
	isRepeated      bool
}

func NewUniqueUniqer() *UniqueUniqer {
	return &UniqueUniqer{}
}

func (u *UniqueUniqer) ShouldSave(line string) bool {
	if !u.hasPreviousLine {
		u.hasPreviousLine = true
		u.previousLine = line
		return false
	}

	if line != u.previousLine {
		u.previousLine = line
		if u.isRepeated {
			u.isRepeated = false
			return false
		} else {
			return true
		}
	}

	if line == u.previousLine {
		u.isRepeated = true
	}

	return false
}

func (u *UniqueUniqer) Finish() bool {
	return !u.isRepeated
}

type RepeatedUniqer struct {
	previousLine    string
	hasPreviousLine bool
	isRepeated      bool
}

func NewRepeatedUniqer() *RepeatedUniqer {
	return &RepeatedUniqer{}
}

func (u *RepeatedUniqer) ShouldSave(line string) bool {
	if !u.hasPreviousLine {
		u.hasPreviousLine = true
		u.previousLine = line
		return false
	}

	if line != u.previousLine {
		u.previousLine = line
		if u.isRepeated {
			u.isRepeated = false
			return true
		} else {
			return false
		}
	}

	if line == u.previousLine {
		u.isRepeated = true
	}

	return false
}

func (u *RepeatedUniqer) Finish() bool {
	return u.isRepeated
}

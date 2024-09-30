package modifier

import (
	"strings"
)

type Modifier interface {
	Modify(line string) string
}

type CountModifier struct {
	currentCount uint
}

func NewCountModifier() *CountModifier {
	return &CountModifier{}
}

func (m *CountModifier) Modify(line string) string {
	m.currentCount++
	return line
}

func (m *CountModifier) Count() uint {
	return m.currentCount
}

func (m *CountModifier) Set(value uint) {
	m.currentCount = value
}

type SkipFieldsModifier struct {
	skipFieldsCount uint
}

func NewSkipFieldsModifier(fieldsCount uint) *SkipFieldsModifier {
	return &SkipFieldsModifier{fieldsCount}
}

func (m *SkipFieldsModifier) Modify(line string) string {
	fields := strings.Fields(line)

	if len(fields) <= int(m.skipFieldsCount) {
		return ""
	}

	return strings.Join(fields[m.skipFieldsCount:], " ")
}

type SkipCharsModifier struct {
	skipCharsCount uint
}

func NewSkipCharsModifier(charsCount uint) *SkipCharsModifier {
	return &SkipCharsModifier{charsCount}
}

func (m *SkipCharsModifier) Modify(line string) string {
	if int(m.skipCharsCount) > len(line) {
		return ""
	}

	return line[m.skipCharsCount:]
}

type IgnoreCaseModifier struct {
}

func NewIgnoreCaseModifier() *IgnoreCaseModifier {
	return &IgnoreCaseModifier{}
}

func (m *IgnoreCaseModifier) Modify(line string) string {
	return strings.ToLower(line)
}

type MultiModifier struct {
	modifiers []Modifier
}

func NewMultiModifier(modifiers []Modifier) Modifier {
	return &MultiModifier{modifiers}
}

func (m *MultiModifier) Modify(line string) string {
	for _, modifier := range m.modifiers {
		line = modifier.Modify(line)
	}

	return line
}

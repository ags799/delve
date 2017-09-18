package terminal

import (
	"github.com/peterh/liner"
	"io"
)

type line interface {
	Prompt(prompt string) (string, error)
	Close() error
	SetCompleter(f liner.Completer)
	ReadHistory(r io.Reader)
	AppendHistory(item string)
	WriteHistory(w io.Writer) (int, error)
}

type defaultLine struct {
	line *liner.State
}

func newLine() line {
	return defaultLine{
		line: liner.NewLiner(),
	}
}

func (l defaultLine) Prompt(prompt string) (string, error) {
	return l.line.Prompt(prompt)
}

func (l defaultLine) Close() error {
	return l.line.Close()
}

func (l defaultLine) SetCompleter(f liner.Completer) {
	l.line.SetCompleter(f)
}

func (l defaultLine) ReadHistory(r io.Reader) {
	l.line.ReadHistory(r)
}

func (l defaultLine) AppendHistory(item string) {
	l.line.AppendHistory(item)
}

func (l defaultLine) WriteHistory(w io.Writer) (int, error) {
	return l.line.WriteHistory(w)
}

type exitLine struct {}

func newExitLine() line {
	return exitLine{}
}

func (l exitLine) Prompt(prompt string) (string, error) {
	return "", nil
}

func (l exitLine) Close() error {
	return nil
}

func (l exitLine) SetCompleter(f liner.Completer) {
	panic("not implemented")
}

func (l exitLine) ReadHistory(r io.Reader) {
	panic("not implemented")
}

func (l exitLine) AppendHistory(item string) {
	panic("not implemented")
}

func (l exitLine) WriteHistory(w io.Writer) (int, error) {
	panic("not implemented")
}

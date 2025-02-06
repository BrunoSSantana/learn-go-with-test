package dependency

import (
	"fmt"
	"io"
)

type GreetWriter interface {
	WriterGreeting(writer io.Writer, name string)
}

type Greet struct{}

func (g *Greet) WriterGreeting(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

type GreetService struct {
	writer GreetWriter
}

func NewGreetService(writer GreetWriter) *GreetService {
	return &GreetService{writer: writer}
}

func (gs *GreetService) Greet(writer io.Writer, name string) {
	gs.writer.WriterGreeting(writer, name)
}

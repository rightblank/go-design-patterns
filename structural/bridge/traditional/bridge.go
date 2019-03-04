package traditional

import (
    "errors"
    "fmt"
    "io"
)

type PrinterAPI interface {
    PrintMessage(s string) error
}

type PrinterImpl1 struct {
}

func (p *PrinterImpl1) PrintMessage(s string) error {
    fmt.Println(s)
    return nil
}

type PrinterImpl2 struct {
    Writer io.Writer
}

func (p *PrinterImpl2) PrintMessage(s string) error {
    if p.Writer == nil {
        return errors.New("You need to pass an io.Writer to PrinterImpl2")
    }
    fmt.Fprintf(p.Writer, "%s", s)
    return nil
}

type NormalPrinter struct {
    Msg     string
    Printer PrinterAPI
}

func (n *NormalPrinter) Print() error {
    n.Printer.PrintMessage(n.Msg)
    return nil
}

type PacktPrinter struct {
    Msg     string
    Printer PrinterAPI
}

func (n *PacktPrinter) Print() error {
    n.Printer.PrintMessage(fmt.Sprintf("Message from Packt: %s", n.Msg))
    return nil
}

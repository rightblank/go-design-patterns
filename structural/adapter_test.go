package structural

import "testing"

func TestPrint(t *testing.T) {
    msg := "Hello world"
    adapter := PrinterAdapter{OldPrinter: &MyLegacyPrinter{}, Msg: msg}
    s := adapter.PrintStored()

    if s != "LegacyPrinter: Adapter: Hello world\n" {
        t.Errorf("message did'nt match %s\n", s)
    }

    adapter = PrinterAdapter{OldPrinter: nil, Msg: msg}
    s = adapter.PrintStored()

    if s != "Hello world" {
        t.Errorf("message did'nt match %s\n", s)
    }
}

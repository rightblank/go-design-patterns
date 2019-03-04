package traditional

import (
    "errors"
    "strings"
    "testing"
)

func TestPrinterImpl1_PrintMessage(t *testing.T) {
    p := PrinterImpl1{}
    err := p.PrintMessage("hello")
    if err != nil {
        t.Errorf("error trying to use API1 implementation: Message: %s\n", err.Error())
    }
}

type TestWriter struct {
    Msg string
}

func (t *TestWriter) Write(p []byte) (int, error) {
    l := len(p)
    if l <= 0 {
        return 0, errors.New("info passed in is empty")
    }
    t.Msg = string(p)
    return l, nil
}

func TestPrinterImpl2_PrintMessage(t *testing.T) {
    p := PrinterImpl2{}
    err := p.PrintMessage("hello")
    if err != nil {
        expectedErrorMsg := "You need to pass an io.Writer to PrinterImpl2"
        if !strings.Contains(err.Error(), expectedErrorMsg) {
            t.Errorf("error message incorrect.\nActual: %s\nExpected: %s\n",
                err.Error(), expectedErrorMsg)
        }

    }

    testWriter := &TestWriter{}
    p = PrinterImpl2{
        Writer: testWriter,
    }
    expectedMsg := "Hello"
    err = p.PrintMessage(expectedMsg)
    if err != nil {
        t.Errorf("Error trying to use the api2 implementation: %s\n", err.Error())
    }

    if testWriter.Msg != expectedMsg {
        t.Fatalf("error message incorrect.\nActual: %s\nExpected: %s\n",
            testWriter.Msg, expectedMsg)
    }
}

func TestNormalPrinter_Print(t *testing.T) {
    expectedMsg := "Hello"
    normalPrinter := NormalPrinter{
        Msg:     expectedMsg,
        Printer: &PrinterImpl1{},
    }
    err := normalPrinter.Print()
    if err != nil {
        t.Errorf(err.Error())
    }

    testWriter := TestWriter{}
    normalPrinter = NormalPrinter{
        Msg: expectedMsg,
        Printer: &PrinterImpl2{
            Writer: &testWriter,
        },
    }

    err = normalPrinter.Print()
    if err != nil {
        t.Error(err.Error())
    }

    if testWriter.Msg != expectedMsg {
        t.Fatalf("error message incorrect.\nActual: %s\nExpected: %s\n",
            testWriter.Msg, expectedMsg)
    }
}

func TestPacktPrinter_Print(t *testing.T) {
    passedMsg := "Hello"
    expectedMsg := "Message from Packt: Hello"
    packtPrinter := PacktPrinter{
        Msg:     passedMsg,
        Printer: &PrinterImpl1{},
    }
    err := packtPrinter.Print()
    if err != nil {
        t.Errorf(err.Error())
    }

    testWriter := TestWriter{}
    packtPrinter = PacktPrinter{
        Msg: passedMsg,
        Printer: &PrinterImpl2{
            Writer: &testWriter,
        },
    }

    err = packtPrinter.Print()
    if err != nil {
        t.Error(err.Error())
    }

    if testWriter.Msg != expectedMsg {
        t.Fatalf("error message incorrect.\nActual: %s\nExpected: %s\n",
            testWriter.Msg, expectedMsg)
    }
}

package main

type Document struct {
}

type Machine interface {
	Print(d Document)
	Fax(d Document)
	Scan(d Document)
}

type MultiFunctionPrinter struct {
}

func (m MultiFunctionPrinter) Print(d Document) {

}

func (m MultiFunctionPrinter) Fax(d Document) {

}

func (m MultiFunctionPrinter) Scan(d Document) {

}

type OldFashionedPrinter struct {
}

func (m OldFashionedPrinter) Print(d Document) {

}

func (m OldFashionedPrinter) Fax(d Document) {

}

func (m OldFashionedPrinter) Scan(d Document) {

}

type Printer interface {
	Print(d Document)
}

type Scanner interface {
	Scan(d Document)
}

type MyPrint struct {
}

func (m MyPrint) Print(d Document) {

}

type Photocopier struct {
}

func (p Photocopier) Scan(d Document) {

}

func (p Photocopier) Print(d Document) {

}

type MultipleFunctionDevice interface {
	Printer
	Scanner
}

// decorator
type MultiFunctionMachine struct {
	printer Printer
	scanner Scanner
}

func (m MultiFunctionMachine) Print(d Document) {
	m.printer.Print(d)
}

func main() {
}

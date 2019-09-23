package typeassert

type Fooer interface {
	String() string
	Foo()
}

type Barer interface {
	String() string
	Bar()
}

type Foo struct{ s string }

type Bar struct{ s string }

func (f Foo) String() string { return f.s }
func (f Foo) Foo()           {}

func (b Bar) String() string { return b.s }
func (b Bar) Bar()           {}

func fooToBar(arg Fooer) Barer {
	bar := arg.(Barer)
	return bar
}

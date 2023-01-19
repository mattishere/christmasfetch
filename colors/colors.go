package colors


type Color struct {
	Name string
	Color string
}

var (
	Blue = Color{Name: "blue", Color: "\x1b[34m"}
	Yellow = Color{Name: "yellow", Color: "\x1b[33m"}
	Green = Color{Name: "green", Color: "\x1b[32m"}
	Magenta = Color{Name: "magenta", Color: "\x1b[35m"}
	Red = Color{Name: "red", Color: "\x1b[31m"}
	Black = Color{Name: "black", Color: "\x1b[30m"}
	White = Color{Name: "white", Color: "\x1b[37m"}
	Cyan = Color{Name: "cyan", Color: "\x1b[36m"}
	Reset = Color{Name: "reset", Color: "\x1b[0m"}

	ColorsList = []Color{
		Blue, Yellow, Green, Magenta, Red, Black, White, Cyan, Reset,
	}
)

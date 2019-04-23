package mod

const hello = "Hey"

func Hello(name string) string {
	if name == "" {
		return hello
	}
	return hello + " " + name
}

package utils

func Assert(ok bool, text string)  {
	if !ok {
		panic(text)
	}
}
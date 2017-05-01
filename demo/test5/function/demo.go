package function

type Books struct {
	AddNum
}

func AddNum(num int) int {
	return num * 10
}

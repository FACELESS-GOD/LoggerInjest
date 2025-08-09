package Subscriber

type SubInt interface {
	Subscriber()
}

type SubStruct struct {
}

func Sub() SubStruct {
	sub := SubStruct{}
	return sub
}

func (Sub *SubStruct) Subscriber() {

}

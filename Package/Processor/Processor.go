package Processor

import "log"

type ProcInt interface {
	Process(Channel string, Payload string)
}

type ProcStruct struct{}

func Proc() ProcStruct {
	proc := ProcStruct{}
	return proc
}

func (Proc *ProcStruct) Process(Channel string, Payload string) {
	log.Println(Channel, Payload)
}

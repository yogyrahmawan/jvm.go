package instructions

import "jvmgo/rtda"

// Compare float
type fcmpg struct {NoOperandsInstruction}
func (self *fcmpg) Execute(thread *rtda.Thread) {
    _fcmp(thread, true)
}

type fcmpl struct {NoOperandsInstruction}
func (self *fcmpl) Execute(thread *rtda.Thread) {
    _fcmp(thread, false)
}

func _fcmp(thread *rtda.Thread, gFlag bool) {
    stack := thread.CurrentFrame().OperandStack()
    v2 := stack.PopFloat()
    v1 := stack.PopFloat()
    if v1 > v2 {
        stack.PushInt(1)
    } else if v1 == v2 {
        stack.PushInt(0)
    } else if v1 < v2 {
        stack.PushInt(-1)
    } else if gFlag {
        stack.PushInt(1)
    } else {
        stack.PushInt(-1)
    }
}

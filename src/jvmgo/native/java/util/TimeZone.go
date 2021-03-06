package util

import (
	. "jvmgo/any"
	"jvmgo/jvm/rtda"
	rtc "jvmgo/jvm/rtda/class"
	"time"
)

func init() {
	_tz(getSystemTimeZoneID, "getSystemTimeZoneID", "(Ljava/lang/String;Ljava/lang/String;)Ljava/lang/String;")
}

func _tz(method Any, name, desc string) {
	rtc.RegisterNativeMethod("java/util/TimeZone", name, desc, method)
}

// private static native String getSystemTimeZoneID(String javaHome, String country);
// (Ljava/lang/String;Ljava/lang/String;)Ljava/lang/String;
func getSystemTimeZoneID(frame *rtda.Frame) {
	//vars := frame.LocalVars()
	//javaHomeObj := vars.GetRef(0)
	//countryObj := vars.GetRef(1)

	// todo
	name, _ := time.Now().Zone()
	zoneID := rtda.NewJString(name)

	stack := frame.OperandStack()
	stack.PushRef(zoneID)
}

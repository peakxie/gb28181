package utils

import (
	"strings"
)

//sip:username@domain -> username(DeviceId)
func GetDeviceId(user string) string {
	ind := strings.Index(user, ":")
	if ind == -1 {
		return ""
	}
	indEnd := strings.Index(user, "@")
	if indEnd == -1 {
		return ""
	}

	if ind >= indEnd {
		return ""
	}

	return user[ind+1 : indEnd]
}

func GetTid(method, branch string) string {
	//在匹配服务端事物的ACK中，创建事务的请求的方法为INVITE。所以ACK消息匹配事物的时候需要注意？？？？
	return method + "_" + branch
}

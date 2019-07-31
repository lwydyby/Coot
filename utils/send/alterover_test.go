package send

import (
	"testing"
)

func TestRunner(t *testing.T){
	//SendAlertOver("测试标题","测试内容")
	SendPushBullet(nil)
}
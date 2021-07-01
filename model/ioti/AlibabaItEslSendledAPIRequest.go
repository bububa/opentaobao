package ioti

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaItEslSendledAPIRequest
厂测LED控制 API请求
alibaba.it.esl.sendled

针对厂测生产的的价签，增加led闪灯的接口，进行led 闪灯测试 */
type AlibabaItEslSendledAPIRequest struct {
	model.Params
	// mac
	_macAp string
	// 0、1、2、3：关蓝绿红
	_type string
}

// New

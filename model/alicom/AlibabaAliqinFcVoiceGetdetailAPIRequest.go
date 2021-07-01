package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAliqinFcVoiceGetdetailAPIRequest
获取呼叫详情 API请求
alibaba.aliqin.fc.voice.getdetail

通过呼叫id获取呼叫相关的数据 */
type AlibabaAliqinFcVoiceGetdetailAPIRequest struct {
	model.Params
	// 呼叫唯一ID
	_callId string
	// 语音通知为:11000000300006, 语音验证码为:11010000138001, IVR为:11000000300005, 点击拨号为:11000000300004, SIP为:11000000300009
	_prodId int64
	// Unix时间戳，会查询这个时间点对应那一天的记录（单位毫秒）
	_queryDate int64
}

// New

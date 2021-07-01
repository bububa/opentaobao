package jst

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoJstSmsMessageDirectBatchsendAPIRequest
聚石塔新短信发送接口 API请求
taobao.jst.sms.message.direct.batchsend

聚石塔所见即所得的短信发送接口 */
type TaobaoJstSmsMessageDirectBatchsendAPIRequest struct {
	model.Params
	// 短信签名
	_signName string
	// 短信中带的短链，如果不传则没有短信效果数据
	_url string
	// 短信模版CODE，必须为全变量模板
	_smsTemplateCode string
	// 短信接收号码，json格式，最多200个号码
	_recNum string
	// 短信内容，如果带${url}则会被入参url替换
	_smsContent string
	// 短信扩展码（JSON字符串数组格式），拓展码个数需要和电话号码个数一致。
	_extendNum string
	// 短信任务code，没有请先创建。
	_taskCode string
	// 一个在taskcode下唯一的随机字符串，对于taskSign相同的请求平台认为是商家的同一次短信发送。
	_taskSign string
}

// New

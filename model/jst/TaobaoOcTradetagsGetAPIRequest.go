package jst

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOcTradetagsGetAPIRequest
根据订单查询订单标签 API请求
taobao.oc.tradetags.get

根据订单查询订单标签。<br/>
返回的tag说明:1为官方标，2为自定义标，3为主站只读标签。<br/>
官方标签和自定义标签请看taobao.oc.tradetag.attach 接口说明<br/>
主站只读标签请看:http://open.taobao.com/doc/detail.htm?id=102865<br/> */
type TaobaoOcTradetagsGetAPIRequest struct {
	model.Params
	// 交易主订单id
	_tid int64
	// 是否查询历史标签
	_history int64
	// 不填，则默认只查询1,2。1为官方标，2为自定义标，3为主站只读标签
	_tagTypes []string
	// 不填，则不做标签名称过滤
	_tagNames []string
}

// New

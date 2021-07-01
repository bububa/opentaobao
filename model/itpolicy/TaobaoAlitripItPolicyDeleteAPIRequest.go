package itpolicy

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAlitripItPolicyDeleteAPIRequest
【国际机票销售规则】单条删除 API请求
taobao.alitrip.it.policy.delete

销售规则删除接口，可以根据taobaoId或outId删除，根据outId删除时，如果outId不唯一，返回失败 */
type TaobaoAlitripItPolicyDeleteAPIRequest struct {
	model.Params
	// 扩展字段
	_extendAttributes string
	// 接入方产品id
	_outId string
	// 淘宝政策id
	_taobaoId int64
}

// New

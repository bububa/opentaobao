package itpolicy

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAlitripItPolicyUpdateAPIRequest
【国际机票销售规则】单条更新 API请求
taobao.alitrip.it.policy.update

销售规则更新接口，可以根据taobaoId或outId修改，outId不唯一时，不能用outId修改。 */
type TaobaoAlitripItPolicyUpdateAPIRequest struct {
	model.Params
	// 扩展字段
	_extendAttributes string
	// 接入方产品id
	_outId string
	// 淘宝政策id
	_taobaoId int64
	// 国际机票销售规则
	_topPolicyDo *TopPolicyDo
}

// New

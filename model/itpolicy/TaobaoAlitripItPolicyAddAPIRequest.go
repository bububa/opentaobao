package itpolicy

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAlitripItPolicyAddAPIRequest
【国际机票销售规则】单条新增 API请求
taobao.alitrip.it.policy.add

销售规则新增，成功返回taobaoId */
type TaobaoAlitripItPolicyAddAPIRequest struct {
	model.Params
	// 扩展字段
	_extendAttributes string
	// 国际机票销售规则
	_topPolicyDo *TopPolicyDo
}

// New

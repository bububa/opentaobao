package itpolicy

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAlitripItPolicyGetAPIRequest
【国际机票销售规则】单条查询 API请求
taobao.alitrip.it.policy.get

通过此接口可以查询单条销售规则的详情，可以根据taobaoId或outId查询，用户outId查询时，如果outId不唯一，只返回最新添加的一条数据。taobaoId为新增成功时候返回的唯一id，outId为新增时的policy_id（产品编号） */
type TaobaoAlitripItPolicyGetAPIRequest struct {
	model.Params
	// 预留扩展字段
	_extendAttributes string
	// 接入方产品编号
	_outId string
	// 淘宝政策id
	_taobaoId int64
}

// New

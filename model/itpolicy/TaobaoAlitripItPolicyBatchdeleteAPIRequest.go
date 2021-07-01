package itpolicy

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAlitripItPolicyBatchdeleteAPIRequest
【国际机票销售规则】批量删除 API请求
taobao.alitrip.it.policy.batchdelete

批量删除销售规则，单次删除最大5万，大于5万时候提示失败，需要缩小删除条件。此接口同步返回任务id，异步执行操作。每个接入方最多同时只能有10个处理中的任务，超过后直接返回失败。 */
type TaobaoAlitripItPolicyBatchdeleteAPIRequest struct {
	model.Params
	// 航司二字码，完整匹配
	_airline string
	// 到达，，完整匹配
	_arrCity string
	// 舱位，，完整匹配
	_cabin string
	// 出发，，完整匹配
	_depCity string
	// 产品id，，完整匹配
	_policyId string
	// 0：未发布 1：已发布 2：已过期。不传的话，默认只能删除未发布和已过期的数据
	_statusList []int64
}

// New

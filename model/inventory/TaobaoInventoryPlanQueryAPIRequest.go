package inventory

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoInventoryPlanQueryAPIRequest 计划库存查询 API请求
// taobao.inventory.plan.query
//
// 计划库存查询
type TaobaoInventoryPlanQueryAPIRequest struct {
	model.Params
	// 计划库存查询入参
	_param *InvUnifyPlanTopQuerys
}

// NewTaobaoInventoryPlanQueryRequest 初始化TaobaoInventoryPlanQueryAPIRequest对象
func NewTaobaoInventoryPlanQueryRequest() *TaobaoInventoryPlanQueryAPIRequest {
	return &TaobaoInventoryPlanQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoInventoryPlanQueryAPIRequest) Reset() {
	r._param = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoInventoryPlanQueryAPIRequest) GetApiMethodName() string {
	return "taobao.inventory.plan.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoInventoryPlanQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoInventoryPlanQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 计划库存查询入参
func (r *TaobaoInventoryPlanQueryAPIRequest) SetParam(_param *InvUnifyPlanTopQuerys) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r TaobaoInventoryPlanQueryAPIRequest) GetParam() *InvUnifyPlanTopQuerys {
	return r._param
}

var poolTaobaoInventoryPlanQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoInventoryPlanQueryRequest()
	},
}

// GetTaobaoInventoryPlanQueryRequest 从 sync.Pool 获取 TaobaoInventoryPlanQueryAPIRequest
func GetTaobaoInventoryPlanQueryAPIRequest() *TaobaoInventoryPlanQueryAPIRequest {
	return poolTaobaoInventoryPlanQueryAPIRequest.Get().(*TaobaoInventoryPlanQueryAPIRequest)
}

// ReleaseTaobaoInventoryPlanQueryAPIRequest 将 TaobaoInventoryPlanQueryAPIRequest 放入 sync.Pool
func ReleaseTaobaoInventoryPlanQueryAPIRequest(v *TaobaoInventoryPlanQueryAPIRequest) {
	v.Reset()
	poolTaobaoInventoryPlanQueryAPIRequest.Put(v)
}

package inventory

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoinventoryplanqueryAPIRequest 计划库存查询 API请求
// taobao.inventory.plan.query
//
// 计划库存查询
type TaobaoinventoryplanqueryAPIRequest struct {
	model.Params
	// 计划库存查询入参
	_param *InvUnifyPlanTopQuerys
}

// NewTaobaoinventoryplanqueryRequest 初始化TaobaoinventoryplanqueryAPIRequest对象
func NewTaobaoinventoryplanqueryRequest() *TaobaoinventoryplanqueryAPIRequest {
	return &TaobaoinventoryplanqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoinventoryplanqueryAPIRequest) GetApiMethodName() string {
	return "taobao.inventory.plan.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoinventoryplanqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoinventoryplanqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 计划库存查询入参
func (r *TaobaoinventoryplanqueryAPIRequest) SetParam(_param *InvUnifyPlanTopQuerys) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r TaobaoinventoryplanqueryAPIRequest) GetParam() *InvUnifyPlanTopQuerys {
	return r._param
}

package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallinventoryqueryforstoreAPIRequest 查询后端商品仓库库存 API请求
// tmall.inventory.query.forstore
//
// 商家查询后端商品仓库库存
type TmallinventoryqueryforstoreAPIRequest struct {
	model.Params
	// 查询列表
	_paramList []InventoryQueryForStoreRequest
}

// NewTmallinventoryqueryforstoreRequest 初始化TmallinventoryqueryforstoreAPIRequest对象
func NewTmallinventoryqueryforstoreRequest() *TmallinventoryqueryforstoreAPIRequest {
	return &TmallinventoryqueryforstoreAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallinventoryqueryforstoreAPIRequest) GetApiMethodName() string {
	return "tmall.inventory.query.forstore"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallinventoryqueryforstoreAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallinventoryqueryforstoreAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamList is ParamList Setter
// 查询列表
func (r *TmallinventoryqueryforstoreAPIRequest) SetParamList(_paramList []InventoryQueryForStoreRequest) error {
	r._paramList = _paramList
	r.Set("param_list", _paramList)
	return nil
}

// GetParamList ParamList Getter
func (r TmallinventoryqueryforstoreAPIRequest) GetParamList() []InventoryQueryForStoreRequest {
	return r._paramList
}

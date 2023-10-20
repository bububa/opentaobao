package eleenterpriserestaurant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaeleenterpriserestaurantmultispecmenuAPIRequest 查询餐厅菜单 API请求
// alibaba.ele.enterprise.restaurant.multispecmenu
//
// 查询餐厅菜单
type AlibabaeleenterpriserestaurantmultispecmenuAPIRequest struct {
	model.Params
	// 餐厅ID
	_erestaurantId string
}

// NewAlibabaeleenterpriserestaurantmultispecmenuRequest 初始化AlibabaeleenterpriserestaurantmultispecmenuAPIRequest对象
func NewAlibabaeleenterpriserestaurantmultispecmenuRequest() *AlibabaeleenterpriserestaurantmultispecmenuAPIRequest {
	return &AlibabaeleenterpriserestaurantmultispecmenuAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaeleenterpriserestaurantmultispecmenuAPIRequest) GetApiMethodName() string {
	return "alibaba.ele.enterprise.restaurant.multispecmenu"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaeleenterpriserestaurantmultispecmenuAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaeleenterpriserestaurantmultispecmenuAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetErestaurantId is ErestaurantId Setter
// 餐厅ID
func (r *AlibabaeleenterpriserestaurantmultispecmenuAPIRequest) SetErestaurantId(_erestaurantId string) error {
	r._erestaurantId = _erestaurantId
	r.Set("erestaurant_id", _erestaurantId)
	return nil
}

// GetErestaurantId ErestaurantId Getter
func (r AlibabaeleenterpriserestaurantmultispecmenuAPIRequest) GetErestaurantId() string {
	return r._erestaurantId
}

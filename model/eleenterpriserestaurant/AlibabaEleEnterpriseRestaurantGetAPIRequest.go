package eleenterpriserestaurant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEleEnterpriseRestaurantGetAPIRequest 查询餐厅信息 API请求
// alibaba.ele.enterprise.restaurant.get
//
// 查询餐厅信息
type AlibabaEleEnterpriseRestaurantGetAPIRequest struct {
	model.Params
	// longitude和latitude用英文逗号分隔
	_geo string
	// 餐厅ID
	_erestaurantId string
}

// NewAlibabaEleEnterpriseRestaurantGetRequest 初始化AlibabaEleEnterpriseRestaurantGetAPIRequest对象
func NewAlibabaEleEnterpriseRestaurantGetRequest() *AlibabaEleEnterpriseRestaurantGetAPIRequest {
	return &AlibabaEleEnterpriseRestaurantGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaEleEnterpriseRestaurantGetAPIRequest) GetApiMethodName() string {
	return "alibaba.ele.enterprise.restaurant.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaEleEnterpriseRestaurantGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaEleEnterpriseRestaurantGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetGeo is Geo Setter
// longitude和latitude用英文逗号分隔
func (r *AlibabaEleEnterpriseRestaurantGetAPIRequest) SetGeo(_geo string) error {
	r._geo = _geo
	r.Set("geo", _geo)
	return nil
}

// GetGeo Geo Getter
func (r AlibabaEleEnterpriseRestaurantGetAPIRequest) GetGeo() string {
	return r._geo
}

// SetErestaurantId is ErestaurantId Setter
// 餐厅ID
func (r *AlibabaEleEnterpriseRestaurantGetAPIRequest) SetErestaurantId(_erestaurantId string) error {
	r._erestaurantId = _erestaurantId
	r.Set("erestaurant_id", _erestaurantId)
	return nil
}

// GetErestaurantId ErestaurantId Getter
func (r AlibabaEleEnterpriseRestaurantGetAPIRequest) GetErestaurantId() string {
	return r._erestaurantId
}

package eleenterpriserestaurant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaeleenterpriserestaurantgetAPIRequest 查询餐厅信息 API请求
// alibaba.ele.enterprise.restaurant.get
//
// 查询餐厅信息
type AlibabaeleenterpriserestaurantgetAPIRequest struct {
	model.Params
	// longitude和latitude用英文逗号分隔
	_geo string
	// 餐厅ID
	_erestaurantId string
}

// NewAlibabaeleenterpriserestaurantgetRequest 初始化AlibabaeleenterpriserestaurantgetAPIRequest对象
func NewAlibabaeleenterpriserestaurantgetRequest() *AlibabaeleenterpriserestaurantgetAPIRequest {
	return &AlibabaeleenterpriserestaurantgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaeleenterpriserestaurantgetAPIRequest) GetApiMethodName() string {
	return "alibaba.ele.enterprise.restaurant.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaeleenterpriserestaurantgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaeleenterpriserestaurantgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetGeo is Geo Setter
// longitude和latitude用英文逗号分隔
func (r *AlibabaeleenterpriserestaurantgetAPIRequest) SetGeo(_geo string) error {
	r._geo = _geo
	r.Set("geo", _geo)
	return nil
}

// GetGeo Geo Getter
func (r AlibabaeleenterpriserestaurantgetAPIRequest) GetGeo() string {
	return r._geo
}

// SetErestaurantId is ErestaurantId Setter
// 餐厅ID
func (r *AlibabaeleenterpriserestaurantgetAPIRequest) SetErestaurantId(_erestaurantId string) error {
	r._erestaurantId = _erestaurantId
	r.Set("erestaurant_id", _erestaurantId)
	return nil
}

// GetErestaurantId ErestaurantId Getter
func (r AlibabaeleenterpriserestaurantgetAPIRequest) GetErestaurantId() string {
	return r._erestaurantId
}

package waybill

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaowaybillprivacysubscriptiongetAPIRequest 隐私面单商家订购查询 API请求
// cainiao.waybill.privacy.subscription.get
//
// ISV查询商家是否订购隐私面单
type CainiaowaybillprivacysubscriptiongetAPIRequest struct {
	model.Params
}

// NewCainiaowaybillprivacysubscriptiongetRequest 初始化CainiaowaybillprivacysubscriptiongetAPIRequest对象
func NewCainiaowaybillprivacysubscriptiongetRequest() *CainiaowaybillprivacysubscriptiongetAPIRequest {
	return &CainiaowaybillprivacysubscriptiongetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaowaybillprivacysubscriptiongetAPIRequest) GetApiMethodName() string {
	return "cainiao.waybill.privacy.subscription.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaowaybillprivacysubscriptiongetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaowaybillprivacysubscriptiongetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

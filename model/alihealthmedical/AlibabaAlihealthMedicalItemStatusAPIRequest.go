package alihealthmedical

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthMedicalItemStatusAPIRequest 商品上下架 API请求
// alibaba.alihealth.medical.item.status
//
// 医生通三方机构平台进行服务商品上下架操作
type AlibabaAlihealthMedicalItemStatusAPIRequest struct {
	model.Params
	// 请求入参
	_shelfrequest *ThirdAgencyUpDownShelfRequest
}

// NewAlibabaAlihealthMedicalItemStatusRequest 初始化AlibabaAlihealthMedicalItemStatusAPIRequest对象
func NewAlibabaAlihealthMedicalItemStatusRequest() *AlibabaAlihealthMedicalItemStatusAPIRequest {
	return &AlibabaAlihealthMedicalItemStatusAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthMedicalItemStatusAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.medical.item.status"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthMedicalItemStatusAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetShelfrequest is Shelfrequest Setter
// 请求入参
func (r *AlibabaAlihealthMedicalItemStatusAPIRequest) SetShelfrequest(_shelfrequest *ThirdAgencyUpDownShelfRequest) error {
	r._shelfrequest = _shelfrequest
	r.Set("shelfrequest", _shelfrequest)
	return nil
}

// GetShelfrequest Shelfrequest Getter
func (r AlibabaAlihealthMedicalItemStatusAPIRequest) GetShelfrequest() *ThirdAgencyUpDownShelfRequest {
	return r._shelfrequest
}

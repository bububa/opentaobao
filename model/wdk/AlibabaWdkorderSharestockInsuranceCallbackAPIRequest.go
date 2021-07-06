package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkorderSharestockInsuranceCallbackAPIRequest 共享库存订单投保后回传保单号 API请求
// alibaba.wdkorder.sharestock.insurance.callback
//
// 共享库存订单投保消息获取
type AlibabaWdkorderSharestockInsuranceCallbackAPIRequest struct {
	model.Params
	// 投保单ID
	_insuranceId string
	// 淘宝交易子订单ID
	_tbSubOrderId int64
}

// NewAlibabaWdkorderSharestockInsuranceCallbackRequest 初始化AlibabaWdkorderSharestockInsuranceCallbackAPIRequest对象
func NewAlibabaWdkorderSharestockInsuranceCallbackRequest() *AlibabaWdkorderSharestockInsuranceCallbackAPIRequest {
	return &AlibabaWdkorderSharestockInsuranceCallbackAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkorderSharestockInsuranceCallbackAPIRequest) GetApiMethodName() string {
	return "alibaba.wdkorder.sharestock.insurance.callback"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkorderSharestockInsuranceCallbackAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetInsuranceId is InsuranceId Setter
// 投保单ID
func (r *AlibabaWdkorderSharestockInsuranceCallbackAPIRequest) SetInsuranceId(_insuranceId string) error {
	r._insuranceId = _insuranceId
	r.Set("insurance_id", _insuranceId)
	return nil
}

// GetInsuranceId InsuranceId Getter
func (r AlibabaWdkorderSharestockInsuranceCallbackAPIRequest) GetInsuranceId() string {
	return r._insuranceId
}

// SetTbSubOrderId is TbSubOrderId Setter
// 淘宝交易子订单ID
func (r *AlibabaWdkorderSharestockInsuranceCallbackAPIRequest) SetTbSubOrderId(_tbSubOrderId int64) error {
	r._tbSubOrderId = _tbSubOrderId
	r.Set("tb_sub_order_id", _tbSubOrderId)
	return nil
}

// GetTbSubOrderId TbSubOrderId Getter
func (r AlibabaWdkorderSharestockInsuranceCallbackAPIRequest) GetTbSubOrderId() int64 {
	return r._tbSubOrderId
}

package wdk

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkorderSharestockInsuranceCallbackAPIRequest) Reset() {
	r._insuranceId = ""
	r._tbSubOrderId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkorderSharestockInsuranceCallbackAPIRequest) GetApiMethodName() string {
	return "alibaba.wdkorder.sharestock.insurance.callback"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkorderSharestockInsuranceCallbackAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkorderSharestockInsuranceCallbackAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaWdkorderSharestockInsuranceCallbackAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkorderSharestockInsuranceCallbackRequest()
	},
}

// GetAlibabaWdkorderSharestockInsuranceCallbackRequest 从 sync.Pool 获取 AlibabaWdkorderSharestockInsuranceCallbackAPIRequest
func GetAlibabaWdkorderSharestockInsuranceCallbackAPIRequest() *AlibabaWdkorderSharestockInsuranceCallbackAPIRequest {
	return poolAlibabaWdkorderSharestockInsuranceCallbackAPIRequest.Get().(*AlibabaWdkorderSharestockInsuranceCallbackAPIRequest)
}

// ReleaseAlibabaWdkorderSharestockInsuranceCallbackAPIRequest 将 AlibabaWdkorderSharestockInsuranceCallbackAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkorderSharestockInsuranceCallbackAPIRequest(v *AlibabaWdkorderSharestockInsuranceCallbackAPIRequest) {
	v.Reset()
	poolAlibabaWdkorderSharestockInsuranceCallbackAPIRequest.Put(v)
}

package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkorderSharestockInsuranceGetorderAPIRequest 共享库存订单投保消息获取 API请求
// alibaba.wdkorder.sharestock.insurance.getorder
//
// 共享库存订单投保消息获取
type AlibabaWdkorderSharestockInsuranceGetorderAPIRequest struct {
	model.Params
	// 淘宝子订单ID
	_tbSubOrderId int64
}

// NewAlibabaWdkorderSharestockInsuranceGetorderRequest 初始化AlibabaWdkorderSharestockInsuranceGetorderAPIRequest对象
func NewAlibabaWdkorderSharestockInsuranceGetorderRequest() *AlibabaWdkorderSharestockInsuranceGetorderAPIRequest {
	return &AlibabaWdkorderSharestockInsuranceGetorderAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkorderSharestockInsuranceGetorderAPIRequest) Reset() {
	r._tbSubOrderId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkorderSharestockInsuranceGetorderAPIRequest) GetApiMethodName() string {
	return "alibaba.wdkorder.sharestock.insurance.getorder"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkorderSharestockInsuranceGetorderAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkorderSharestockInsuranceGetorderAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTbSubOrderId is TbSubOrderId Setter
// 淘宝子订单ID
func (r *AlibabaWdkorderSharestockInsuranceGetorderAPIRequest) SetTbSubOrderId(_tbSubOrderId int64) error {
	r._tbSubOrderId = _tbSubOrderId
	r.Set("tb_sub_order_id", _tbSubOrderId)
	return nil
}

// GetTbSubOrderId TbSubOrderId Getter
func (r AlibabaWdkorderSharestockInsuranceGetorderAPIRequest) GetTbSubOrderId() int64 {
	return r._tbSubOrderId
}

var poolAlibabaWdkorderSharestockInsuranceGetorderAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkorderSharestockInsuranceGetorderRequest()
	},
}

// GetAlibabaWdkorderSharestockInsuranceGetorderRequest 从 sync.Pool 获取 AlibabaWdkorderSharestockInsuranceGetorderAPIRequest
func GetAlibabaWdkorderSharestockInsuranceGetorderAPIRequest() *AlibabaWdkorderSharestockInsuranceGetorderAPIRequest {
	return poolAlibabaWdkorderSharestockInsuranceGetorderAPIRequest.Get().(*AlibabaWdkorderSharestockInsuranceGetorderAPIRequest)
}

// ReleaseAlibabaWdkorderSharestockInsuranceGetorderAPIRequest 将 AlibabaWdkorderSharestockInsuranceGetorderAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkorderSharestockInsuranceGetorderAPIRequest(v *AlibabaWdkorderSharestockInsuranceGetorderAPIRequest) {
	v.Reset()
	poolAlibabaWdkorderSharestockInsuranceGetorderAPIRequest.Put(v)
}

package logistic

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEleFengniaoMerchantContractCancelAPIRequest 蜂鸟商户解约接口 API请求
// alibaba.ele.fengniao.merchant.contract.cancel
//
// 通过调用此接口，商家及商家下的所有门店解除蜂鸟物流服务
type AlibabaEleFengniaoMerchantContractCancelAPIRequest struct {
	model.Params
	// 系统自动生成
	_param *Param
}

// NewAlibabaEleFengniaoMerchantContractCancelRequest 初始化AlibabaEleFengniaoMerchantContractCancelAPIRequest对象
func NewAlibabaEleFengniaoMerchantContractCancelRequest() *AlibabaEleFengniaoMerchantContractCancelAPIRequest {
	return &AlibabaEleFengniaoMerchantContractCancelAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaEleFengniaoMerchantContractCancelAPIRequest) Reset() {
	r._param = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaEleFengniaoMerchantContractCancelAPIRequest) GetApiMethodName() string {
	return "alibaba.ele.fengniao.merchant.contract.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaEleFengniaoMerchantContractCancelAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaEleFengniaoMerchantContractCancelAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 系统自动生成
func (r *AlibabaEleFengniaoMerchantContractCancelAPIRequest) SetParam(_param *Param) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaEleFengniaoMerchantContractCancelAPIRequest) GetParam() *Param {
	return r._param
}

var poolAlibabaEleFengniaoMerchantContractCancelAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaEleFengniaoMerchantContractCancelRequest()
	},
}

// GetAlibabaEleFengniaoMerchantContractCancelRequest 从 sync.Pool 获取 AlibabaEleFengniaoMerchantContractCancelAPIRequest
func GetAlibabaEleFengniaoMerchantContractCancelAPIRequest() *AlibabaEleFengniaoMerchantContractCancelAPIRequest {
	return poolAlibabaEleFengniaoMerchantContractCancelAPIRequest.Get().(*AlibabaEleFengniaoMerchantContractCancelAPIRequest)
}

// ReleaseAlibabaEleFengniaoMerchantContractCancelAPIRequest 将 AlibabaEleFengniaoMerchantContractCancelAPIRequest 放入 sync.Pool
func ReleaseAlibabaEleFengniaoMerchantContractCancelAPIRequest(v *AlibabaEleFengniaoMerchantContractCancelAPIRequest) {
	v.Reset()
	poolAlibabaEleFengniaoMerchantContractCancelAPIRequest.Put(v)
}

package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaelefengniaomerchantcontractcancelAPIRequest 蜂鸟商户解约接口 API请求
// alibaba.ele.fengniao.merchant.contract.cancel
//
// 通过调用此接口，商家及商家下的所有门店解除蜂鸟物流服务
type AlibabaelefengniaomerchantcontractcancelAPIRequest struct {
	model.Params
	// 系统自动生成
	_param *Param
}

// NewAlibabaelefengniaomerchantcontractcancelRequest 初始化AlibabaelefengniaomerchantcontractcancelAPIRequest对象
func NewAlibabaelefengniaomerchantcontractcancelRequest() *AlibabaelefengniaomerchantcontractcancelAPIRequest {
	return &AlibabaelefengniaomerchantcontractcancelAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaelefengniaomerchantcontractcancelAPIRequest) GetApiMethodName() string {
	return "alibaba.ele.fengniao.merchant.contract.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaelefengniaomerchantcontractcancelAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaelefengniaomerchantcontractcancelAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 系统自动生成
func (r *AlibabaelefengniaomerchantcontractcancelAPIRequest) SetParam(_param *Param) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaelefengniaomerchantcontractcancelAPIRequest) GetParam() *Param {
	return r._param
}

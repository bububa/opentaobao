package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaelefengniaochainstorecontractcancelAPIRequest 门店解约接口 API请求
// alibaba.ele.fengniao.chainstore.contract.cancel
//
// 调用成功后，门店和蜂鸟解除物流合同，不能再使用此门店推单
type AlibabaelefengniaochainstorecontractcancelAPIRequest struct {
	model.Params
	// 系统自动生成
	_param *AlibabaelefengniaochainstorecontractcancelData
}

// NewAlibabaelefengniaochainstorecontractcancelRequest 初始化AlibabaelefengniaochainstorecontractcancelAPIRequest对象
func NewAlibabaelefengniaochainstorecontractcancelRequest() *AlibabaelefengniaochainstorecontractcancelAPIRequest {
	return &AlibabaelefengniaochainstorecontractcancelAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaelefengniaochainstorecontractcancelAPIRequest) GetApiMethodName() string {
	return "alibaba.ele.fengniao.chainstore.contract.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaelefengniaochainstorecontractcancelAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaelefengniaochainstorecontractcancelAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 系统自动生成
func (r *AlibabaelefengniaochainstorecontractcancelAPIRequest) SetParam(_param *AlibabaelefengniaochainstorecontractcancelData) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaelefengniaochainstorecontractcancelAPIRequest) GetParam() *AlibabaelefengniaochainstorecontractcancelData {
	return r._param
}

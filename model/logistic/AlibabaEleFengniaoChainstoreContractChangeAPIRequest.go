package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaelefengniaochainstorecontractchangeAPIRequest 门店改签合同接口 API请求
// alibaba.ele.fengniao.chainstore.contract.change
//
// 通过调用接口，门店改签物流服务包
type AlibabaelefengniaochainstorecontractchangeAPIRequest struct {
	model.Params
	// 系统自动生成
	_param *Param
}

// NewAlibabaelefengniaochainstorecontractchangeRequest 初始化AlibabaelefengniaochainstorecontractchangeAPIRequest对象
func NewAlibabaelefengniaochainstorecontractchangeRequest() *AlibabaelefengniaochainstorecontractchangeAPIRequest {
	return &AlibabaelefengniaochainstorecontractchangeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaelefengniaochainstorecontractchangeAPIRequest) GetApiMethodName() string {
	return "alibaba.ele.fengniao.chainstore.contract.change"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaelefengniaochainstorecontractchangeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaelefengniaochainstorecontractchangeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 系统自动生成
func (r *AlibabaelefengniaochainstorecontractchangeAPIRequest) SetParam(_param *Param) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaelefengniaochainstorecontractchangeAPIRequest) GetParam() *Param {
	return r._param
}

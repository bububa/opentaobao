package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEleFengniaoChainstoreContractChangeAPIRequest 门店改签合同接口 API请求
// alibaba.ele.fengniao.chainstore.contract.change
//
// 通过调用接口，门店改签物流服务包
type AlibabaEleFengniaoChainstoreContractChangeAPIRequest struct {
	model.Params
	// 系统自动生成
	_param *Param
}

// NewAlibabaEleFengniaoChainstoreContractChangeRequest 初始化AlibabaEleFengniaoChainstoreContractChangeAPIRequest对象
func NewAlibabaEleFengniaoChainstoreContractChangeRequest() *AlibabaEleFengniaoChainstoreContractChangeAPIRequest {
	return &AlibabaEleFengniaoChainstoreContractChangeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaEleFengniaoChainstoreContractChangeAPIRequest) GetApiMethodName() string {
	return "alibaba.ele.fengniao.chainstore.contract.change"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaEleFengniaoChainstoreContractChangeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaEleFengniaoChainstoreContractChangeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 系统自动生成
func (r *AlibabaEleFengniaoChainstoreContractChangeAPIRequest) SetParam(_param *Param) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaEleFengniaoChainstoreContractChangeAPIRequest) GetParam() *Param {
	return r._param
}

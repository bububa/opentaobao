package logistic

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEleFengniaoChainstoreContractCancelAPIRequest 门店解约接口 API请求
// alibaba.ele.fengniao.chainstore.contract.cancel
//
// 调用成功后，门店和蜂鸟解除物流合同，不能再使用此门店推单
type AlibabaEleFengniaoChainstoreContractCancelAPIRequest struct {
	model.Params
	// 系统自动生成
	_param *AlibabaEleFengniaoChainstoreContractCancelData
}

// NewAlibabaEleFengniaoChainstoreContractCancelRequest 初始化AlibabaEleFengniaoChainstoreContractCancelAPIRequest对象
func NewAlibabaEleFengniaoChainstoreContractCancelRequest() *AlibabaEleFengniaoChainstoreContractCancelAPIRequest {
	return &AlibabaEleFengniaoChainstoreContractCancelAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaEleFengniaoChainstoreContractCancelAPIRequest) Reset() {
	r._param = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaEleFengniaoChainstoreContractCancelAPIRequest) GetApiMethodName() string {
	return "alibaba.ele.fengniao.chainstore.contract.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaEleFengniaoChainstoreContractCancelAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaEleFengniaoChainstoreContractCancelAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 系统自动生成
func (r *AlibabaEleFengniaoChainstoreContractCancelAPIRequest) SetParam(_param *AlibabaEleFengniaoChainstoreContractCancelData) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaEleFengniaoChainstoreContractCancelAPIRequest) GetParam() *AlibabaEleFengniaoChainstoreContractCancelData {
	return r._param
}

var poolAlibabaEleFengniaoChainstoreContractCancelAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaEleFengniaoChainstoreContractCancelRequest()
	},
}

// GetAlibabaEleFengniaoChainstoreContractCancelRequest 从 sync.Pool 获取 AlibabaEleFengniaoChainstoreContractCancelAPIRequest
func GetAlibabaEleFengniaoChainstoreContractCancelAPIRequest() *AlibabaEleFengniaoChainstoreContractCancelAPIRequest {
	return poolAlibabaEleFengniaoChainstoreContractCancelAPIRequest.Get().(*AlibabaEleFengniaoChainstoreContractCancelAPIRequest)
}

// ReleaseAlibabaEleFengniaoChainstoreContractCancelAPIRequest 将 AlibabaEleFengniaoChainstoreContractCancelAPIRequest 放入 sync.Pool
func ReleaseAlibabaEleFengniaoChainstoreContractCancelAPIRequest(v *AlibabaEleFengniaoChainstoreContractCancelAPIRequest) {
	v.Reset()
	poolAlibabaEleFengniaoChainstoreContractCancelAPIRequest.Put(v)
}

package logistic

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEleFengniaoChainstoreUpdateAPIRequest 修改门店信息接口 API请求
// alibaba.ele.fengniao.chainstore.update
//
// 修改门店的经纬度，文本地址，电话，门店名
type AlibabaEleFengniaoChainstoreUpdateAPIRequest struct {
	model.Params
	// 入参
	_param *Param
}

// NewAlibabaEleFengniaoChainstoreUpdateRequest 初始化AlibabaEleFengniaoChainstoreUpdateAPIRequest对象
func NewAlibabaEleFengniaoChainstoreUpdateRequest() *AlibabaEleFengniaoChainstoreUpdateAPIRequest {
	return &AlibabaEleFengniaoChainstoreUpdateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaEleFengniaoChainstoreUpdateAPIRequest) Reset() {
	r._param = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaEleFengniaoChainstoreUpdateAPIRequest) GetApiMethodName() string {
	return "alibaba.ele.fengniao.chainstore.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaEleFengniaoChainstoreUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaEleFengniaoChainstoreUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 入参
func (r *AlibabaEleFengniaoChainstoreUpdateAPIRequest) SetParam(_param *Param) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaEleFengniaoChainstoreUpdateAPIRequest) GetParam() *Param {
	return r._param
}

var poolAlibabaEleFengniaoChainstoreUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaEleFengniaoChainstoreUpdateRequest()
	},
}

// GetAlibabaEleFengniaoChainstoreUpdateRequest 从 sync.Pool 获取 AlibabaEleFengniaoChainstoreUpdateAPIRequest
func GetAlibabaEleFengniaoChainstoreUpdateAPIRequest() *AlibabaEleFengniaoChainstoreUpdateAPIRequest {
	return poolAlibabaEleFengniaoChainstoreUpdateAPIRequest.Get().(*AlibabaEleFengniaoChainstoreUpdateAPIRequest)
}

// ReleaseAlibabaEleFengniaoChainstoreUpdateAPIRequest 将 AlibabaEleFengniaoChainstoreUpdateAPIRequest 放入 sync.Pool
func ReleaseAlibabaEleFengniaoChainstoreUpdateAPIRequest(v *AlibabaEleFengniaoChainstoreUpdateAPIRequest) {
	v.Reset()
	poolAlibabaEleFengniaoChainstoreUpdateAPIRequest.Put(v)
}

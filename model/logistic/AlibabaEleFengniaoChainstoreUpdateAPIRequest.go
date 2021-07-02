package logistic

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaEleFengniaoChainstoreUpdateAPIRequest) GetApiMethodName() string {
	return "alibaba.ele.fengniao.chainstore.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaEleFengniaoChainstoreUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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

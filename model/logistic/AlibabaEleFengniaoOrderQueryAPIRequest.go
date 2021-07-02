package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEleFengniaoOrderQueryAPIRequest 查询订单基本信息 API请求
// alibaba.ele.fengniao.order.query
//
// 查询订单基本信息
type AlibabaEleFengniaoOrderQueryAPIRequest struct {
	model.Params
	// 参数
	_param *Param
}

// NewAlibabaEleFengniaoOrderQueryRequest 初始化AlibabaEleFengniaoOrderQueryAPIRequest对象
func NewAlibabaEleFengniaoOrderQueryRequest() *AlibabaEleFengniaoOrderQueryAPIRequest {
	return &AlibabaEleFengniaoOrderQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaEleFengniaoOrderQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.ele.fengniao.order.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaEleFengniaoOrderQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParam is Param Setter
// 参数
func (r *AlibabaEleFengniaoOrderQueryAPIRequest) SetParam(_param *Param) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaEleFengniaoOrderQueryAPIRequest) GetParam() *Param {
	return r._param
}

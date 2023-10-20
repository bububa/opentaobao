package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaelefengniaoorderqueryAPIRequest 查询订单基本信息 API请求
// alibaba.ele.fengniao.order.query
//
// 查询订单基本信息
type AlibabaelefengniaoorderqueryAPIRequest struct {
	model.Params
	// 参数
	_param *Param
}

// NewAlibabaelefengniaoorderqueryRequest 初始化AlibabaelefengniaoorderqueryAPIRequest对象
func NewAlibabaelefengniaoorderqueryRequest() *AlibabaelefengniaoorderqueryAPIRequest {
	return &AlibabaelefengniaoorderqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaelefengniaoorderqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.ele.fengniao.order.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaelefengniaoorderqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaelefengniaoorderqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 参数
func (r *AlibabaelefengniaoorderqueryAPIRequest) SetParam(_param *Param) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaelefengniaoorderqueryAPIRequest) GetParam() *Param {
	return r._param
}

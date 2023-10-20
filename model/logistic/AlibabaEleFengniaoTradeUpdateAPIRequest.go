package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaelefengniaotradeupdateAPIRequest 更新蜂鸟扣费状态 API请求
// alibaba.ele.fengniao.trade.update
//
// 汇金扣费成功后，回调该接口更新扣费状态
type AlibabaelefengniaotradeupdateAPIRequest struct {
	model.Params
	// param 参数
	_param *Param
}

// NewAlibabaelefengniaotradeupdateRequest 初始化AlibabaelefengniaotradeupdateAPIRequest对象
func NewAlibabaelefengniaotradeupdateRequest() *AlibabaelefengniaotradeupdateAPIRequest {
	return &AlibabaelefengniaotradeupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaelefengniaotradeupdateAPIRequest) GetApiMethodName() string {
	return "alibaba.ele.fengniao.trade.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaelefengniaotradeupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaelefengniaotradeupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// param 参数
func (r *AlibabaelefengniaotradeupdateAPIRequest) SetParam(_param *Param) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaelefengniaotradeupdateAPIRequest) GetParam() *Param {
	return r._param
}

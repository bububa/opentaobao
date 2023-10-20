package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaelefengniaochainstoreupdateAPIRequest 修改门店信息接口 API请求
// alibaba.ele.fengniao.chainstore.update
//
// 修改门店的经纬度，文本地址，电话，门店名
type AlibabaelefengniaochainstoreupdateAPIRequest struct {
	model.Params
	// 入参
	_param *Param
}

// NewAlibabaelefengniaochainstoreupdateRequest 初始化AlibabaelefengniaochainstoreupdateAPIRequest对象
func NewAlibabaelefengniaochainstoreupdateRequest() *AlibabaelefengniaochainstoreupdateAPIRequest {
	return &AlibabaelefengniaochainstoreupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaelefengniaochainstoreupdateAPIRequest) GetApiMethodName() string {
	return "alibaba.ele.fengniao.chainstore.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaelefengniaochainstoreupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaelefengniaochainstoreupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 入参
func (r *AlibabaelefengniaochainstoreupdateAPIRequest) SetParam(_param *Param) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaelefengniaochainstoreupdateAPIRequest) GetParam() *Param {
	return r._param
}

package sungari

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosungaridisposequeryAPIRequest 商品商家处置结果查询 API请求
// taobao.sungari.dispose.query
//
// 红盾云桥同政府合作，将线下寄函的商品商家处置转为线上处理
type TaobaosungaridisposequeryAPIRequest struct {
	model.Params
	// 查询的key列表
	_paramList []string
}

// NewTaobaosungaridisposequeryRequest 初始化TaobaosungaridisposequeryAPIRequest对象
func NewTaobaosungaridisposequeryRequest() *TaobaosungaridisposequeryAPIRequest {
	return &TaobaosungaridisposequeryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosungaridisposequeryAPIRequest) GetApiMethodName() string {
	return "taobao.sungari.dispose.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosungaridisposequeryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosungaridisposequeryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamList is ParamList Setter
// 查询的key列表
func (r *TaobaosungaridisposequeryAPIRequest) SetParamList(_paramList []string) error {
	r._paramList = _paramList
	r.Set("param_list", _paramList)
	return nil
}

// GetParamList ParamList Getter
func (r TaobaosungaridisposequeryAPIRequest) GetParamList() []string {
	return r._paramList
}

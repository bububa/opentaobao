package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoomniorderdtdqueryAPIRequest 门店自送根据核销码查订单 API请求
// taobao.omniorder.dtd.query
//
// 门店自送根据核销码码查询订单信息
type TaobaoomniorderdtdqueryAPIRequest struct {
	model.Params
	// 核销码
	_code string
}

// NewTaobaoomniorderdtdqueryRequest 初始化TaobaoomniorderdtdqueryAPIRequest对象
func NewTaobaoomniorderdtdqueryRequest() *TaobaoomniorderdtdqueryAPIRequest {
	return &TaobaoomniorderdtdqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoomniorderdtdqueryAPIRequest) GetApiMethodName() string {
	return "taobao.omniorder.dtd.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoomniorderdtdqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoomniorderdtdqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCode is Code Setter
// 核销码
func (r *TaobaoomniorderdtdqueryAPIRequest) SetCode(_code string) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// GetCode Code Getter
func (r TaobaoomniorderdtdqueryAPIRequest) GetCode() string {
	return r._code
}

package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenChannelinventoryQueryAPIRequest 渠道库存查询接口 API请求
// taobao.qimen.channelinventory.query
//
// 渠道库存查询
type TaobaoQimenChannelinventoryQueryAPIRequest struct {
	model.Params
	//
	_request *RequestDo
}

// NewTaobaoQimenChannelinventoryQueryRequest 初始化TaobaoQimenChannelinventoryQueryAPIRequest对象
func NewTaobaoQimenChannelinventoryQueryRequest() *TaobaoQimenChannelinventoryQueryAPIRequest {
	return &TaobaoQimenChannelinventoryQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQimenChannelinventoryQueryAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.channelinventory.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQimenChannelinventoryQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Request Setter
//
func (r *TaobaoQimenChannelinventoryQueryAPIRequest) SetRequest(_request *RequestDo) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// Get Request Getter
func (r TaobaoQimenChannelinventoryQueryAPIRequest) GetRequest() *RequestDo {
	return r._request
}

package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenInventorybatchQueryAPIRequest 商品单仓批次库存查询接口 API请求
// taobao.qimen.inventorybatch.query
//
// ERP 通过该接口查询指定商品的单仓批次库存
type TaobaoQimenInventorybatchQueryAPIRequest struct {
	model.Params
	// request
	_request *TaobaoQimenInventorybatchQueryRequest
}

// NewTaobaoQimenInventorybatchQueryRequest 初始化TaobaoQimenInventorybatchQueryAPIRequest对象
func NewTaobaoQimenInventorybatchQueryRequest() *TaobaoQimenInventorybatchQueryAPIRequest {
	return &TaobaoQimenInventorybatchQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQimenInventorybatchQueryAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.inventorybatch.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQimenInventorybatchQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Request Setter
// request
func (r *TaobaoQimenInventorybatchQueryAPIRequest) SetRequest(_request *TaobaoQimenInventorybatchQueryRequest) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// Get Request Getter
func (r TaobaoQimenInventorybatchQueryAPIRequest) GetRequest() *TaobaoQimenInventorybatchQueryRequest {
	return r._request
}

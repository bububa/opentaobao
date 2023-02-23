package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenTransferorderCreateAPIRequest 调拨单创建 API请求
// taobao.qimen.transferorder.create
//
// 调拨单创建
type TaobaoQimenTransferorderCreateAPIRequest struct {
	model.Params
	//
	_request *TaobaoQimenTransferorderCreateStruct
}

// NewTaobaoQimenTransferorderCreateRequest 初始化TaobaoQimenTransferorderCreateAPIRequest对象
func NewTaobaoQimenTransferorderCreateRequest() *TaobaoQimenTransferorderCreateAPIRequest {
	return &TaobaoQimenTransferorderCreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQimenTransferorderCreateAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.transferorder.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQimenTransferorderCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoQimenTransferorderCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest is Request Setter
func (r *TaobaoQimenTransferorderCreateAPIRequest) SetRequest(_request *TaobaoQimenTransferorderCreateStruct) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r TaobaoQimenTransferorderCreateAPIRequest) GetRequest() *TaobaoQimenTransferorderCreateStruct {
	return r._request
}

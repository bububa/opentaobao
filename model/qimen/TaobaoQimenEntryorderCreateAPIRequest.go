package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenEntryorderCreateAPIRequest 入库单创建接口 API请求
// taobao.qimen.entryorder.create
//
// ERP调用接口，创建入库单;
type TaobaoQimenEntryorderCreateAPIRequest struct {
	model.Params
	//
	_request *EntryOrderCreateRequest
}

// NewTaobaoQimenEntryorderCreateRequest 初始化TaobaoQimenEntryorderCreateAPIRequest对象
func NewTaobaoQimenEntryorderCreateRequest() *TaobaoQimenEntryorderCreateAPIRequest {
	return &TaobaoQimenEntryorderCreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQimenEntryorderCreateAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.entryorder.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQimenEntryorderCreateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Request Setter
//
func (r *TaobaoQimenEntryorderCreateAPIRequest) SetRequest(_request *EntryOrderCreateRequest) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// Get Request Getter
func (r TaobaoQimenEntryorderCreateAPIRequest) GetRequest() *EntryOrderCreateRequest {
	return r._request
}

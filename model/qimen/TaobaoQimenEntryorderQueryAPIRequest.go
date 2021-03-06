package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenEntryorderQueryAPIRequest 入库单查询接口 API请求
// taobao.qimen.entryorder.query
//
// ERP调用接口，查询入库单信息;
type TaobaoQimenEntryorderQueryAPIRequest struct {
	model.Params
	//
	_request *EntryOrderQueryRequest
}

// NewTaobaoQimenEntryorderQueryRequest 初始化TaobaoQimenEntryorderQueryAPIRequest对象
func NewTaobaoQimenEntryorderQueryRequest() *TaobaoQimenEntryorderQueryAPIRequest {
	return &TaobaoQimenEntryorderQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQimenEntryorderQueryAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.entryorder.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQimenEntryorderQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetRequest is Request Setter
//
func (r *TaobaoQimenEntryorderQueryAPIRequest) SetRequest(_request *EntryOrderQueryRequest) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r TaobaoQimenEntryorderQueryAPIRequest) GetRequest() *EntryOrderQueryRequest {
	return r._request
}

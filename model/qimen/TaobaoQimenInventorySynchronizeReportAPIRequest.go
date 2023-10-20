package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqimeninventorysynchronizereportAPIRequest 库存状态同步确认接口 API请求
// taobao.qimen.inventory.synchronize.report
//
// 库存状态同步确认接口
type TaobaoqimeninventorysynchronizereportAPIRequest struct {
	model.Params
	//
	_request *TaobaoqimeninventorysynchronizereportRequest
}

// NewTaobaoqimeninventorysynchronizereportRequest 初始化TaobaoqimeninventorysynchronizereportAPIRequest对象
func NewTaobaoqimeninventorysynchronizereportRequest() *TaobaoqimeninventorysynchronizereportAPIRequest {
	return &TaobaoqimeninventorysynchronizereportAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoqimeninventorysynchronizereportAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.inventory.synchronize.report"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoqimeninventorysynchronizereportAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoqimeninventorysynchronizereportAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest is Request Setter
func (r *TaobaoqimeninventorysynchronizereportAPIRequest) SetRequest(_request *TaobaoqimeninventorysynchronizereportRequest) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r TaobaoqimeninventorysynchronizereportAPIRequest) GetRequest() *TaobaoqimeninventorysynchronizereportRequest {
	return r._request
}

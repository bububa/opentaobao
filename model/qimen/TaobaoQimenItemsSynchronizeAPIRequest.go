package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoQimenItemsSynchronizeAPIRequest
商品同步接口 (批量) API请求
taobao.qimen.items.synchronize

ERP调用奇门的接口,批量同步商品信息给WMS */
type TaobaoQimenItemsSynchronizeAPIRequest struct {
	model.Params
	//
	_request *ItemsSynchronizeRequest
}

// NewTaobaoQimenItemsSynchronizeRequest 初始化TaobaoQimenItemsSynchronizeAPIRequest对象
func NewTaobaoQimenItemsSynchronizeRequest() *TaobaoQimenItemsSynchronizeAPIRequest {
	return &TaobaoQimenItemsSynchronizeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQimenItemsSynchronizeAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.items.synchronize"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQimenItemsSynchronizeAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Request Setter
//
func (r *TaobaoQimenItemsSynchronizeAPIRequest) SetRequest(_request *ItemsSynchronizeRequest) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// Get Request Getter
func (r TaobaoQimenItemsSynchronizeAPIRequest) GetRequest() *ItemsSynchronizeRequest {
	return r._request
}

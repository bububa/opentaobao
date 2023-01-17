package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenSingleitemSynchronizeAPIRequest 商品同步接口 API请求
// taobao.qimen.singleitem.synchronize
//
// ERP调用奇门的接口,同步商品信息给WMS
type TaobaoQimenSingleitemSynchronizeAPIRequest struct {
	model.Params
	//
	_request *ItemSynRequest
}

// NewTaobaoQimenSingleitemSynchronizeRequest 初始化TaobaoQimenSingleitemSynchronizeAPIRequest对象
func NewTaobaoQimenSingleitemSynchronizeRequest() *TaobaoQimenSingleitemSynchronizeAPIRequest {
	return &TaobaoQimenSingleitemSynchronizeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQimenSingleitemSynchronizeAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.singleitem.synchronize"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQimenSingleitemSynchronizeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoQimenSingleitemSynchronizeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest is Request Setter
func (r *TaobaoQimenSingleitemSynchronizeAPIRequest) SetRequest(_request *ItemSynRequest) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r TaobaoQimenSingleitemSynchronizeAPIRequest) GetRequest() *ItemSynRequest {
	return r._request
}

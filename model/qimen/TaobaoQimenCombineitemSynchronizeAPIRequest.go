package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqimencombineitemsynchronizeAPIRequest 组合商品接口 API请求
// taobao.qimen.combineitem.synchronize
//
// ERP调用奇门的接口,将商品信息同步给WMS
type TaobaoqimencombineitemsynchronizeAPIRequest struct {
	model.Params
	//
	_request *CombineItemSyncRequest
}

// NewTaobaoqimencombineitemsynchronizeRequest 初始化TaobaoqimencombineitemsynchronizeAPIRequest对象
func NewTaobaoqimencombineitemsynchronizeRequest() *TaobaoqimencombineitemsynchronizeAPIRequest {
	return &TaobaoqimencombineitemsynchronizeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoqimencombineitemsynchronizeAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.combineitem.synchronize"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoqimencombineitemsynchronizeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoqimencombineitemsynchronizeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest is Request Setter
func (r *TaobaoqimencombineitemsynchronizeAPIRequest) SetRequest(_request *CombineItemSyncRequest) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r TaobaoqimencombineitemsynchronizeAPIRequest) GetRequest() *CombineItemSyncRequest {
	return r._request
}

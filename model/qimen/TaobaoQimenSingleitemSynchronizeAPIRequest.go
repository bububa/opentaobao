package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqimensingleitemsynchronizeAPIRequest 商品同步接口 API请求
// taobao.qimen.singleitem.synchronize
//
// taobao.qimen.singleitem.synchronize
type TaobaoqimensingleitemsynchronizeAPIRequest struct {
	model.Params
	//
	_request *ItemSynRequest
}

// NewTaobaoqimensingleitemsynchronizeRequest 初始化TaobaoqimensingleitemsynchronizeAPIRequest对象
func NewTaobaoqimensingleitemsynchronizeRequest() *TaobaoqimensingleitemsynchronizeAPIRequest {
	return &TaobaoqimensingleitemsynchronizeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoqimensingleitemsynchronizeAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.singleitem.synchronize"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoqimensingleitemsynchronizeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoqimensingleitemsynchronizeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest is Request Setter
func (r *TaobaoqimensingleitemsynchronizeAPIRequest) SetRequest(_request *ItemSynRequest) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r TaobaoqimensingleitemsynchronizeAPIRequest) GetRequest() *ItemSynRequest {
	return r._request
}

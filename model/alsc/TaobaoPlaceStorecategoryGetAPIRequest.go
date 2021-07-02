package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPlaceStorecategoryGetAPIRequest 获取门店类目信息 API请求
// taobao.place.storecategory.get
//
// 获取门店类目信息
type TaobaoPlaceStorecategoryGetAPIRequest struct {
	model.Params
}

// NewTaobaoPlaceStorecategoryGetRequest 初始化TaobaoPlaceStorecategoryGetAPIRequest对象
func NewTaobaoPlaceStorecategoryGetRequest() *TaobaoPlaceStorecategoryGetAPIRequest {
	return &TaobaoPlaceStorecategoryGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoPlaceStorecategoryGetAPIRequest) GetApiMethodName() string {
	return "taobao.place.storecategory.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoPlaceStorecategoryGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

package util

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTimeGetAPIRequest 获取淘宝系统当前时间 API请求
// taobao.time.get
//
// 获取淘宝系统当前时间
type TaobaoTimeGetAPIRequest struct {
	model.Params
}

// NewTaobaoTimeGetRequest 初始化TaobaoTimeGetAPIRequest对象
func NewTaobaoTimeGetRequest() *TaobaoTimeGetAPIRequest {
	return &TaobaoTimeGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTimeGetAPIRequest) GetApiMethodName() string {
	return "taobao.time.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTimeGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

package interact

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabainteractshopfavorAPIRequest 店铺收藏 API请求
// alibaba.interact.shop.favor
//
// 店铺收藏mtop接口开放鉴权接口，无入参出参，无安全风险，mtop接口开发 酒仙。
type AlibabainteractshopfavorAPIRequest struct {
	model.Params
}

// NewAlibabainteractshopfavorRequest 初始化AlibabainteractshopfavorAPIRequest对象
func NewAlibabainteractshopfavorRequest() *AlibabainteractshopfavorAPIRequest {
	return &AlibabainteractshopfavorAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabainteractshopfavorAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.shop.favor"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabainteractshopfavorAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabainteractshopfavorAPIRequest) GetRawParams() model.Params {
	return r.Params
}

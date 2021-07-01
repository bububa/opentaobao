package interact

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaInteractShopFavorAPIRequest
店铺收藏 API请求
alibaba.interact.shop.favor

店铺收藏mtop接口开放鉴权接口，无入参出参，无安全风险，mtop接口开发 酒仙。 */
type AlibabaInteractShopFavorAPIRequest struct {
	model.Params
}

// NewAlibabaInteractShopFavorRequest 初始化AlibabaInteractShopFavorAPIRequest对象
func NewAlibabaInteractShopFavorRequest() *AlibabaInteractShopFavorAPIRequest {
	return &AlibabaInteractShopFavorAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaInteractShopFavorAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.shop.favor"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaInteractShopFavorAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

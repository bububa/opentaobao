package interact

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractCoinBuyerAddAPIRequest 平台向买家发放淘金币 API请求
// alibaba.interact.coin.buyer.add
//
// 手淘开放专用接口，没有数据返回，仅用于手淘容器中jssdk接口鉴权。ISV调用该接口向买家发放平台淘金币，需要优惠平台运营审核ISV资质。
type AlibabaInteractCoinBuyerAddAPIRequest struct {
	model.Params
}

// NewAlibabaInteractCoinBuyerAddRequest 初始化AlibabaInteractCoinBuyerAddAPIRequest对象
func NewAlibabaInteractCoinBuyerAddRequest() *AlibabaInteractCoinBuyerAddAPIRequest {
	return &AlibabaInteractCoinBuyerAddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaInteractCoinBuyerAddAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.coin.buyer.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaInteractCoinBuyerAddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaInteractCoinBuyerAddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

package interact

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractIsvlotteryIsvdrawAPIRequest 天猫奖池鉴权接口 API请求
// alibaba.interact.isvlottery.isvdraw
//
// 鉴权接口，为tida.isvdraw接口鉴权
type AlibabaInteractIsvlotteryIsvdrawAPIRequest struct {
	model.Params
}

// NewAlibabaInteractIsvlotteryIsvdrawRequest 初始化AlibabaInteractIsvlotteryIsvdrawAPIRequest对象
func NewAlibabaInteractIsvlotteryIsvdrawRequest() *AlibabaInteractIsvlotteryIsvdrawAPIRequest {
	return &AlibabaInteractIsvlotteryIsvdrawAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaInteractIsvlotteryIsvdrawAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.isvlottery.isvdraw"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaInteractIsvlotteryIsvdrawAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

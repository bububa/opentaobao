package interact

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractShopFavorAPIRequest 店铺收藏 API请求
// alibaba.interact.shop.favor
//
// 店铺收藏mtop接口开放鉴权接口，无入参出参，无安全风险，mtop接口开发 酒仙。
type AlibabaInteractShopFavorAPIRequest struct {
	model.Params
}

// NewAlibabaInteractShopFavorRequest 初始化AlibabaInteractShopFavorAPIRequest对象
func NewAlibabaInteractShopFavorRequest() *AlibabaInteractShopFavorAPIRequest {
	return &AlibabaInteractShopFavorAPIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaInteractShopFavorAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaInteractShopFavorAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.shop.favor"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaInteractShopFavorAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaInteractShopFavorAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolAlibabaInteractShopFavorAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaInteractShopFavorRequest()
	},
}

// GetAlibabaInteractShopFavorRequest 从 sync.Pool 获取 AlibabaInteractShopFavorAPIRequest
func GetAlibabaInteractShopFavorAPIRequest() *AlibabaInteractShopFavorAPIRequest {
	return poolAlibabaInteractShopFavorAPIRequest.Get().(*AlibabaInteractShopFavorAPIRequest)
}

// ReleaseAlibabaInteractShopFavorAPIRequest 将 AlibabaInteractShopFavorAPIRequest 放入 sync.Pool
func ReleaseAlibabaInteractShopFavorAPIRequest(v *AlibabaInteractShopFavorAPIRequest) {
	v.Reset()
	poolAlibabaInteractShopFavorAPIRequest.Put(v)
}

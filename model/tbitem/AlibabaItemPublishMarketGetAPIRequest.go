package tbitem

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaItemPublishMarketGetAPIRequest 获取商家可发布商品的市场信息 API请求
// alibaba.item.publish.market.get
//
// 获取商家可发布商品的市场信息
type AlibabaItemPublishMarketGetAPIRequest struct {
	model.Params
}

// NewAlibabaItemPublishMarketGetRequest 初始化AlibabaItemPublishMarketGetAPIRequest对象
func NewAlibabaItemPublishMarketGetRequest() *AlibabaItemPublishMarketGetAPIRequest {
	return &AlibabaItemPublishMarketGetAPIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaItemPublishMarketGetAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaItemPublishMarketGetAPIRequest) GetApiMethodName() string {
	return "alibaba.item.publish.market.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaItemPublishMarketGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaItemPublishMarketGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolAlibabaItemPublishMarketGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaItemPublishMarketGetRequest()
	},
}

// GetAlibabaItemPublishMarketGetRequest 从 sync.Pool 获取 AlibabaItemPublishMarketGetAPIRequest
func GetAlibabaItemPublishMarketGetAPIRequest() *AlibabaItemPublishMarketGetAPIRequest {
	return poolAlibabaItemPublishMarketGetAPIRequest.Get().(*AlibabaItemPublishMarketGetAPIRequest)
}

// ReleaseAlibabaItemPublishMarketGetAPIRequest 将 AlibabaItemPublishMarketGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaItemPublishMarketGetAPIRequest(v *AlibabaItemPublishMarketGetAPIRequest) {
	v.Reset()
	poolAlibabaItemPublishMarketGetAPIRequest.Put(v)
}

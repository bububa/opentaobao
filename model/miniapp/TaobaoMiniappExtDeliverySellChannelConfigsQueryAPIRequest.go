package miniapp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoMiniappExtDeliverySellChannelConfigsQueryAPIRequest 查询商家配置的信息 API请求
// taobao.miniapp.ext.delivery.sell.channel.configs.query
//
// 查询商家配置的信息
type TaobaoMiniappExtDeliverySellChannelConfigsQueryAPIRequest struct {
	model.Params
	// 入参
	_req *SellerChannelQueryRequest
}

// NewTaobaoMiniappExtDeliverySellChannelConfigsQueryRequest 初始化TaobaoMiniappExtDeliverySellChannelConfigsQueryAPIRequest对象
func NewTaobaoMiniappExtDeliverySellChannelConfigsQueryRequest() *TaobaoMiniappExtDeliverySellChannelConfigsQueryAPIRequest {
	return &TaobaoMiniappExtDeliverySellChannelConfigsQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoMiniappExtDeliverySellChannelConfigsQueryAPIRequest) Reset() {
	r._req = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoMiniappExtDeliverySellChannelConfigsQueryAPIRequest) GetApiMethodName() string {
	return "taobao.miniapp.ext.delivery.sell.channel.configs.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoMiniappExtDeliverySellChannelConfigsQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoMiniappExtDeliverySellChannelConfigsQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetReq is Req Setter
// 入参
func (r *TaobaoMiniappExtDeliverySellChannelConfigsQueryAPIRequest) SetReq(_req *SellerChannelQueryRequest) error {
	r._req = _req
	r.Set("req", _req)
	return nil
}

// GetReq Req Getter
func (r TaobaoMiniappExtDeliverySellChannelConfigsQueryAPIRequest) GetReq() *SellerChannelQueryRequest {
	return r._req
}

var poolTaobaoMiniappExtDeliverySellChannelConfigsQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoMiniappExtDeliverySellChannelConfigsQueryRequest()
	},
}

// GetTaobaoMiniappExtDeliverySellChannelConfigsQueryRequest 从 sync.Pool 获取 TaobaoMiniappExtDeliverySellChannelConfigsQueryAPIRequest
func GetTaobaoMiniappExtDeliverySellChannelConfigsQueryAPIRequest() *TaobaoMiniappExtDeliverySellChannelConfigsQueryAPIRequest {
	return poolTaobaoMiniappExtDeliverySellChannelConfigsQueryAPIRequest.Get().(*TaobaoMiniappExtDeliverySellChannelConfigsQueryAPIRequest)
}

// ReleaseTaobaoMiniappExtDeliverySellChannelConfigsQueryAPIRequest 将 TaobaoMiniappExtDeliverySellChannelConfigsQueryAPIRequest 放入 sync.Pool
func ReleaseTaobaoMiniappExtDeliverySellChannelConfigsQueryAPIRequest(v *TaobaoMiniappExtDeliverySellChannelConfigsQueryAPIRequest) {
	v.Reset()
	poolTaobaoMiniappExtDeliverySellChannelConfigsQueryAPIRequest.Put(v)
}

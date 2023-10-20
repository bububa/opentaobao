package shop

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSellercatsListGetAPIRequest 获取前台展示的店铺内卖家自定义商品类目 API请求
// taobao.sellercats.list.get
//
// 此API获取当前卖家店铺在淘宝前端被展示的浏览导航类目（面向买家）
type TaobaoSellercatsListGetAPIRequest struct {
	model.Params
}

// NewTaobaoSellercatsListGetRequest 初始化TaobaoSellercatsListGetAPIRequest对象
func NewTaobaoSellercatsListGetRequest() *TaobaoSellercatsListGetAPIRequest {
	return &TaobaoSellercatsListGetAPIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoSellercatsListGetAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSellercatsListGetAPIRequest) GetApiMethodName() string {
	return "taobao.sellercats.list.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSellercatsListGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSellercatsListGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolTaobaoSellercatsListGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoSellercatsListGetRequest()
	},
}

// GetTaobaoSellercatsListGetRequest 从 sync.Pool 获取 TaobaoSellercatsListGetAPIRequest
func GetTaobaoSellercatsListGetAPIRequest() *TaobaoSellercatsListGetAPIRequest {
	return poolTaobaoSellercatsListGetAPIRequest.Get().(*TaobaoSellercatsListGetAPIRequest)
}

// ReleaseTaobaoSellercatsListGetAPIRequest 将 TaobaoSellercatsListGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoSellercatsListGetAPIRequest(v *TaobaoSellercatsListGetAPIRequest) {
	v.Reset()
	poolTaobaoSellercatsListGetAPIRequest.Put(v)
}

package caipiao

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoCaipiaoGoodsInfoGetAPIRequest 根据卖家id与appkey获取商品信息 API请求
// taobao.caipiao.goods.info.get
//
// 根据卖家id与appkey获取商品信息。
type TaobaoCaipiaoGoodsInfoGetAPIRequest struct {
	model.Params
}

// NewTaobaoCaipiaoGoodsInfoGetRequest 初始化TaobaoCaipiaoGoodsInfoGetAPIRequest对象
func NewTaobaoCaipiaoGoodsInfoGetRequest() *TaobaoCaipiaoGoodsInfoGetAPIRequest {
	return &TaobaoCaipiaoGoodsInfoGetAPIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoCaipiaoGoodsInfoGetAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoCaipiaoGoodsInfoGetAPIRequest) GetApiMethodName() string {
	return "taobao.caipiao.goods.info.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoCaipiaoGoodsInfoGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoCaipiaoGoodsInfoGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolTaobaoCaipiaoGoodsInfoGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoCaipiaoGoodsInfoGetRequest()
	},
}

// GetTaobaoCaipiaoGoodsInfoGetRequest 从 sync.Pool 获取 TaobaoCaipiaoGoodsInfoGetAPIRequest
func GetTaobaoCaipiaoGoodsInfoGetAPIRequest() *TaobaoCaipiaoGoodsInfoGetAPIRequest {
	return poolTaobaoCaipiaoGoodsInfoGetAPIRequest.Get().(*TaobaoCaipiaoGoodsInfoGetAPIRequest)
}

// ReleaseTaobaoCaipiaoGoodsInfoGetAPIRequest 将 TaobaoCaipiaoGoodsInfoGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoCaipiaoGoodsInfoGetAPIRequest(v *TaobaoCaipiaoGoodsInfoGetAPIRequest) {
	v.Reset()
	poolTaobaoCaipiaoGoodsInfoGetAPIRequest.Put(v)
}

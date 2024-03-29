package travel

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelBaseinfoCruiseGetAPIRequest 【API3.0】度假线路商品发布时基础信息获取接口：邮轮扩展信息获取 API请求
// taobao.alitrip.travel.baseinfo.cruise.get
//
// 旅行度假新商品发布时可用的扩展接口，用于获取邮轮类目相关扩展信息。
type TaobaoAlitripTravelBaseinfoCruiseGetAPIRequest struct {
	model.Params
	// true-获取国际邮轮类目扩展信息；false-获取国内邮轮类目扩展信息
	_isOverseas bool
}

// NewTaobaoAlitripTravelBaseinfoCruiseGetRequest 初始化TaobaoAlitripTravelBaseinfoCruiseGetAPIRequest对象
func NewTaobaoAlitripTravelBaseinfoCruiseGetRequest() *TaobaoAlitripTravelBaseinfoCruiseGetAPIRequest {
	return &TaobaoAlitripTravelBaseinfoCruiseGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoAlitripTravelBaseinfoCruiseGetAPIRequest) Reset() {
	r._isOverseas = false
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripTravelBaseinfoCruiseGetAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.travel.baseinfo.cruise.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripTravelBaseinfoCruiseGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAlitripTravelBaseinfoCruiseGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetIsOverseas is IsOverseas Setter
// true-获取国际邮轮类目扩展信息；false-获取国内邮轮类目扩展信息
func (r *TaobaoAlitripTravelBaseinfoCruiseGetAPIRequest) SetIsOverseas(_isOverseas bool) error {
	r._isOverseas = _isOverseas
	r.Set("is_overseas", _isOverseas)
	return nil
}

// GetIsOverseas IsOverseas Getter
func (r TaobaoAlitripTravelBaseinfoCruiseGetAPIRequest) GetIsOverseas() bool {
	return r._isOverseas
}

var poolTaobaoAlitripTravelBaseinfoCruiseGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoAlitripTravelBaseinfoCruiseGetRequest()
	},
}

// GetTaobaoAlitripTravelBaseinfoCruiseGetRequest 从 sync.Pool 获取 TaobaoAlitripTravelBaseinfoCruiseGetAPIRequest
func GetTaobaoAlitripTravelBaseinfoCruiseGetAPIRequest() *TaobaoAlitripTravelBaseinfoCruiseGetAPIRequest {
	return poolTaobaoAlitripTravelBaseinfoCruiseGetAPIRequest.Get().(*TaobaoAlitripTravelBaseinfoCruiseGetAPIRequest)
}

// ReleaseTaobaoAlitripTravelBaseinfoCruiseGetAPIRequest 将 TaobaoAlitripTravelBaseinfoCruiseGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoAlitripTravelBaseinfoCruiseGetAPIRequest(v *TaobaoAlitripTravelBaseinfoCruiseGetAPIRequest) {
	v.Reset()
	poolTaobaoAlitripTravelBaseinfoCruiseGetAPIRequest.Put(v)
}

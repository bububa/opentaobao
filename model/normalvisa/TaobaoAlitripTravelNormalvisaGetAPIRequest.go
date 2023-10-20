package normalvisa

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelNormalvisaGetAPIRequest 获取签证记录 API请求
// taobao.alitrip.travel.normalvisa.get
//
// 用于获取普通签证的记录信息
type TaobaoAlitripTravelNormalvisaGetAPIRequest struct {
	model.Params
	// 订单号
	_bizOrderId int64
}

// NewTaobaoAlitripTravelNormalvisaGetRequest 初始化TaobaoAlitripTravelNormalvisaGetAPIRequest对象
func NewTaobaoAlitripTravelNormalvisaGetRequest() *TaobaoAlitripTravelNormalvisaGetAPIRequest {
	return &TaobaoAlitripTravelNormalvisaGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoAlitripTravelNormalvisaGetAPIRequest) Reset() {
	r._bizOrderId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripTravelNormalvisaGetAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.travel.normalvisa.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripTravelNormalvisaGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAlitripTravelNormalvisaGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBizOrderId is BizOrderId Setter
// 订单号
func (r *TaobaoAlitripTravelNormalvisaGetAPIRequest) SetBizOrderId(_bizOrderId int64) error {
	r._bizOrderId = _bizOrderId
	r.Set("biz_order_id", _bizOrderId)
	return nil
}

// GetBizOrderId BizOrderId Getter
func (r TaobaoAlitripTravelNormalvisaGetAPIRequest) GetBizOrderId() int64 {
	return r._bizOrderId
}

var poolTaobaoAlitripTravelNormalvisaGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoAlitripTravelNormalvisaGetRequest()
	},
}

// GetTaobaoAlitripTravelNormalvisaGetRequest 从 sync.Pool 获取 TaobaoAlitripTravelNormalvisaGetAPIRequest
func GetTaobaoAlitripTravelNormalvisaGetAPIRequest() *TaobaoAlitripTravelNormalvisaGetAPIRequest {
	return poolTaobaoAlitripTravelNormalvisaGetAPIRequest.Get().(*TaobaoAlitripTravelNormalvisaGetAPIRequest)
}

// ReleaseTaobaoAlitripTravelNormalvisaGetAPIRequest 将 TaobaoAlitripTravelNormalvisaGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoAlitripTravelNormalvisaGetAPIRequest(v *TaobaoAlitripTravelNormalvisaGetAPIRequest) {
	v.Reset()
	poolTaobaoAlitripTravelNormalvisaGetAPIRequest.Put(v)
}

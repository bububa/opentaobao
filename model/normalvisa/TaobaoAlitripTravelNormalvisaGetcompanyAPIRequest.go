package normalvisa

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelNormalvisaGetcompanyAPIRequest 获取物流公司信息 API请求
// taobao.alitrip.travel.normalvisa.getcompany
//
// 获取物流公司信息
type TaobaoAlitripTravelNormalvisaGetcompanyAPIRequest struct {
	model.Params
	// true：取5个重要的物流公司 false：取所有的物流公司
	_param0 bool
}

// NewTaobaoAlitripTravelNormalvisaGetcompanyRequest 初始化TaobaoAlitripTravelNormalvisaGetcompanyAPIRequest对象
func NewTaobaoAlitripTravelNormalvisaGetcompanyRequest() *TaobaoAlitripTravelNormalvisaGetcompanyAPIRequest {
	return &TaobaoAlitripTravelNormalvisaGetcompanyAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoAlitripTravelNormalvisaGetcompanyAPIRequest) Reset() {
	r._param0 = false
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripTravelNormalvisaGetcompanyAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.travel.normalvisa.getcompany"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripTravelNormalvisaGetcompanyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAlitripTravelNormalvisaGetcompanyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// true：取5个重要的物流公司 false：取所有的物流公司
func (r *TaobaoAlitripTravelNormalvisaGetcompanyAPIRequest) SetParam0(_param0 bool) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r TaobaoAlitripTravelNormalvisaGetcompanyAPIRequest) GetParam0() bool {
	return r._param0
}

var poolTaobaoAlitripTravelNormalvisaGetcompanyAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoAlitripTravelNormalvisaGetcompanyRequest()
	},
}

// GetTaobaoAlitripTravelNormalvisaGetcompanyRequest 从 sync.Pool 获取 TaobaoAlitripTravelNormalvisaGetcompanyAPIRequest
func GetTaobaoAlitripTravelNormalvisaGetcompanyAPIRequest() *TaobaoAlitripTravelNormalvisaGetcompanyAPIRequest {
	return poolTaobaoAlitripTravelNormalvisaGetcompanyAPIRequest.Get().(*TaobaoAlitripTravelNormalvisaGetcompanyAPIRequest)
}

// ReleaseTaobaoAlitripTravelNormalvisaGetcompanyAPIRequest 将 TaobaoAlitripTravelNormalvisaGetcompanyAPIRequest 放入 sync.Pool
func ReleaseTaobaoAlitripTravelNormalvisaGetcompanyAPIRequest(v *TaobaoAlitripTravelNormalvisaGetcompanyAPIRequest) {
	v.Reset()
	poolTaobaoAlitripTravelNormalvisaGetcompanyAPIRequest.Put(v)
}

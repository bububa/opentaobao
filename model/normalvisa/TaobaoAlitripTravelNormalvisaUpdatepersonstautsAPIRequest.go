package normalvisa

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelNormalvisaUpdatepersonstautsAPIRequest 更新签证办理进度 API请求
// taobao.alitrip.travel.normalvisa.updatepersonstauts
//
// 更新签证办理进度
type TaobaoAlitripTravelNormalvisaUpdatepersonstautsAPIRequest struct {
	model.Params
	// 更新信息
	_normalVisaUpdateUnitList []NormalVisaUpdateUnit
	// 订单号
	_bizOrderId int64
}

// NewTaobaoAlitripTravelNormalvisaUpdatepersonstautsRequest 初始化TaobaoAlitripTravelNormalvisaUpdatepersonstautsAPIRequest对象
func NewTaobaoAlitripTravelNormalvisaUpdatepersonstautsRequest() *TaobaoAlitripTravelNormalvisaUpdatepersonstautsAPIRequest {
	return &TaobaoAlitripTravelNormalvisaUpdatepersonstautsAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoAlitripTravelNormalvisaUpdatepersonstautsAPIRequest) Reset() {
	r._normalVisaUpdateUnitList = r._normalVisaUpdateUnitList[:0]
	r._bizOrderId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripTravelNormalvisaUpdatepersonstautsAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.travel.normalvisa.updatepersonstauts"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripTravelNormalvisaUpdatepersonstautsAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAlitripTravelNormalvisaUpdatepersonstautsAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNormalVisaUpdateUnitList is NormalVisaUpdateUnitList Setter
// 更新信息
func (r *TaobaoAlitripTravelNormalvisaUpdatepersonstautsAPIRequest) SetNormalVisaUpdateUnitList(_normalVisaUpdateUnitList []NormalVisaUpdateUnit) error {
	r._normalVisaUpdateUnitList = _normalVisaUpdateUnitList
	r.Set("normal_visa_update_unit_list", _normalVisaUpdateUnitList)
	return nil
}

// GetNormalVisaUpdateUnitList NormalVisaUpdateUnitList Getter
func (r TaobaoAlitripTravelNormalvisaUpdatepersonstautsAPIRequest) GetNormalVisaUpdateUnitList() []NormalVisaUpdateUnit {
	return r._normalVisaUpdateUnitList
}

// SetBizOrderId is BizOrderId Setter
// 订单号
func (r *TaobaoAlitripTravelNormalvisaUpdatepersonstautsAPIRequest) SetBizOrderId(_bizOrderId int64) error {
	r._bizOrderId = _bizOrderId
	r.Set("biz_order_id", _bizOrderId)
	return nil
}

// GetBizOrderId BizOrderId Getter
func (r TaobaoAlitripTravelNormalvisaUpdatepersonstautsAPIRequest) GetBizOrderId() int64 {
	return r._bizOrderId
}

var poolTaobaoAlitripTravelNormalvisaUpdatepersonstautsAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoAlitripTravelNormalvisaUpdatepersonstautsRequest()
	},
}

// GetTaobaoAlitripTravelNormalvisaUpdatepersonstautsRequest 从 sync.Pool 获取 TaobaoAlitripTravelNormalvisaUpdatepersonstautsAPIRequest
func GetTaobaoAlitripTravelNormalvisaUpdatepersonstautsAPIRequest() *TaobaoAlitripTravelNormalvisaUpdatepersonstautsAPIRequest {
	return poolTaobaoAlitripTravelNormalvisaUpdatepersonstautsAPIRequest.Get().(*TaobaoAlitripTravelNormalvisaUpdatepersonstautsAPIRequest)
}

// ReleaseTaobaoAlitripTravelNormalvisaUpdatepersonstautsAPIRequest 将 TaobaoAlitripTravelNormalvisaUpdatepersonstautsAPIRequest 放入 sync.Pool
func ReleaseTaobaoAlitripTravelNormalvisaUpdatepersonstautsAPIRequest(v *TaobaoAlitripTravelNormalvisaUpdatepersonstautsAPIRequest) {
	v.Reset()
	poolTaobaoAlitripTravelNormalvisaUpdatepersonstautsAPIRequest.Put(v)
}

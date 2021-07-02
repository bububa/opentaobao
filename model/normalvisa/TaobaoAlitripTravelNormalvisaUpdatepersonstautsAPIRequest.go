package normalvisa

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripTravelNormalvisaUpdatepersonstautsAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.travel.normalvisa.updatepersonstauts"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripTravelNormalvisaUpdatepersonstautsAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is NormalVisaUpdateUnitList Setter
// 更新信息
func (r *TaobaoAlitripTravelNormalvisaUpdatepersonstautsAPIRequest) SetNormalVisaUpdateUnitList(_normalVisaUpdateUnitList []NormalVisaUpdateUnit) error {
	r._normalVisaUpdateUnitList = _normalVisaUpdateUnitList
	r.Set("normal_visa_update_unit_list", _normalVisaUpdateUnitList)
	return nil
}

// Get NormalVisaUpdateUnitList Getter
func (r TaobaoAlitripTravelNormalvisaUpdatepersonstautsAPIRequest) GetNormalVisaUpdateUnitList() []NormalVisaUpdateUnit {
	return r._normalVisaUpdateUnitList
}

// Set is BizOrderId Setter
// 订单号
func (r *TaobaoAlitripTravelNormalvisaUpdatepersonstautsAPIRequest) SetBizOrderId(_bizOrderId int64) error {
	r._bizOrderId = _bizOrderId
	r.Set("biz_order_id", _bizOrderId)
	return nil
}

// Get BizOrderId Getter
func (r TaobaoAlitripTravelNormalvisaUpdatepersonstautsAPIRequest) GetBizOrderId() int64 {
	return r._bizOrderId
}

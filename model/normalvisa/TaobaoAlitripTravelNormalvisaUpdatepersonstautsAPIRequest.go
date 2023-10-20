package normalvisa

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitriptravelnormalvisaupdatepersonstautsAPIRequest 更新签证办理进度 API请求
// taobao.alitrip.travel.normalvisa.updatepersonstauts
//
// 更新签证办理进度
type TaobaoalitriptravelnormalvisaupdatepersonstautsAPIRequest struct {
	model.Params
	// 更新信息
	_normalVisaUpdateUnitList []NormalVisaUpdateUnit
	// 订单号
	_bizOrderId int64
}

// NewTaobaoalitriptravelnormalvisaupdatepersonstautsRequest 初始化TaobaoalitriptravelnormalvisaupdatepersonstautsAPIRequest对象
func NewTaobaoalitriptravelnormalvisaupdatepersonstautsRequest() *TaobaoalitriptravelnormalvisaupdatepersonstautsAPIRequest {
	return &TaobaoalitriptravelnormalvisaupdatepersonstautsAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoalitriptravelnormalvisaupdatepersonstautsAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.travel.normalvisa.updatepersonstauts"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoalitriptravelnormalvisaupdatepersonstautsAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoalitriptravelnormalvisaupdatepersonstautsAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNormalVisaUpdateUnitList is NormalVisaUpdateUnitList Setter
// 更新信息
func (r *TaobaoalitriptravelnormalvisaupdatepersonstautsAPIRequest) SetNormalVisaUpdateUnitList(_normalVisaUpdateUnitList []NormalVisaUpdateUnit) error {
	r._normalVisaUpdateUnitList = _normalVisaUpdateUnitList
	r.Set("normal_visa_update_unit_list", _normalVisaUpdateUnitList)
	return nil
}

// GetNormalVisaUpdateUnitList NormalVisaUpdateUnitList Getter
func (r TaobaoalitriptravelnormalvisaupdatepersonstautsAPIRequest) GetNormalVisaUpdateUnitList() []NormalVisaUpdateUnit {
	return r._normalVisaUpdateUnitList
}

// SetBizOrderId is BizOrderId Setter
// 订单号
func (r *TaobaoalitriptravelnormalvisaupdatepersonstautsAPIRequest) SetBizOrderId(_bizOrderId int64) error {
	r._bizOrderId = _bizOrderId
	r.Set("biz_order_id", _bizOrderId)
	return nil
}

// GetBizOrderId BizOrderId Getter
func (r TaobaoalitriptravelnormalvisaupdatepersonstautsAPIRequest) GetBizOrderId() int64 {
	return r._bizOrderId
}

package normalvisa

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitriptravelnormalvisastoreuserAPIRequest 代填办理人信息 API请求
// taobao.alitrip.travel.normalvisa.storeuser
//
// 卖家代填买家填写办理人信息
type TaobaoalitriptravelnormalvisastoreuserAPIRequest struct {
	model.Params
	// 列表：签证人信息列表
	_normalVisaUserUnitList []NormalVisaUserUnit
	// 订单id
	_bizOrderId int64
}

// NewTaobaoalitriptravelnormalvisastoreuserRequest 初始化TaobaoalitriptravelnormalvisastoreuserAPIRequest对象
func NewTaobaoalitriptravelnormalvisastoreuserRequest() *TaobaoalitriptravelnormalvisastoreuserAPIRequest {
	return &TaobaoalitriptravelnormalvisastoreuserAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoalitriptravelnormalvisastoreuserAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.travel.normalvisa.storeuser"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoalitriptravelnormalvisastoreuserAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoalitriptravelnormalvisastoreuserAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNormalVisaUserUnitList is NormalVisaUserUnitList Setter
// 列表：签证人信息列表
func (r *TaobaoalitriptravelnormalvisastoreuserAPIRequest) SetNormalVisaUserUnitList(_normalVisaUserUnitList []NormalVisaUserUnit) error {
	r._normalVisaUserUnitList = _normalVisaUserUnitList
	r.Set("normal_visa_user_unit_list", _normalVisaUserUnitList)
	return nil
}

// GetNormalVisaUserUnitList NormalVisaUserUnitList Getter
func (r TaobaoalitriptravelnormalvisastoreuserAPIRequest) GetNormalVisaUserUnitList() []NormalVisaUserUnit {
	return r._normalVisaUserUnitList
}

// SetBizOrderId is BizOrderId Setter
// 订单id
func (r *TaobaoalitriptravelnormalvisastoreuserAPIRequest) SetBizOrderId(_bizOrderId int64) error {
	r._bizOrderId = _bizOrderId
	r.Set("biz_order_id", _bizOrderId)
	return nil
}

// GetBizOrderId BizOrderId Getter
func (r TaobaoalitriptravelnormalvisastoreuserAPIRequest) GetBizOrderId() int64 {
	return r._bizOrderId
}

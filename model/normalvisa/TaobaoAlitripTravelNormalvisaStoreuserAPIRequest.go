package normalvisa

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAlitripTravelNormalvisaStoreuserAPIRequest
代填办理人信息 API请求
taobao.alitrip.travel.normalvisa.storeuser

卖家代填买家填写办理人信息 */
type TaobaoAlitripTravelNormalvisaStoreuserAPIRequest struct {
	model.Params
	// 订单id
	_bizOrderId int64
	// 列表：签证人信息列表
	_normalVisaUserUnitList []NormalVisaUserUnit
}

// NewTaobaoAlitripTravelNormalvisaStoreuserRequest 初始化TaobaoAlitripTravelNormalvisaStoreuserAPIRequest对象
func NewTaobaoAlitripTravelNormalvisaStoreuserRequest() *TaobaoAlitripTravelNormalvisaStoreuserAPIRequest {
	return &TaobaoAlitripTravelNormalvisaStoreuserAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripTravelNormalvisaStoreuserAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.travel.normalvisa.storeuser"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripTravelNormalvisaStoreuserAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is BizOrderId Setter
// 订单id
func (r *TaobaoAlitripTravelNormalvisaStoreuserAPIRequest) SetBizOrderId(_bizOrderId int64) error {
	r._bizOrderId = _bizOrderId
	r.Set("biz_order_id", _bizOrderId)
	return nil
}

// Get BizOrderId Getter
func (r TaobaoAlitripTravelNormalvisaStoreuserAPIRequest) GetBizOrderId() int64 {
	return r._bizOrderId
}

// Set is NormalVisaUserUnitList Setter
// 列表：签证人信息列表
func (r *TaobaoAlitripTravelNormalvisaStoreuserAPIRequest) SetNormalVisaUserUnitList(_normalVisaUserUnitList []NormalVisaUserUnit) error {
	r._normalVisaUserUnitList = _normalVisaUserUnitList
	r.Set("normal_visa_user_unit_list", _normalVisaUserUnitList)
	return nil
}

// Get NormalVisaUserUnitList Getter
func (r TaobaoAlitripTravelNormalvisaStoreuserAPIRequest) GetNormalVisaUserUnitList() []NormalVisaUserUnit {
	return r._normalVisaUserUnitList
}

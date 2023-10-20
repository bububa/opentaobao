package normalvisa

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelNormalvisaStoreuserAPIRequest 代填办理人信息 API请求
// taobao.alitrip.travel.normalvisa.storeuser
//
// 卖家代填买家填写办理人信息
type TaobaoAlitripTravelNormalvisaStoreuserAPIRequest struct {
	model.Params
	// 列表：签证人信息列表
	_normalVisaUserUnitList []NormalVisaUserUnit
	// 订单id
	_bizOrderId int64
}

// NewTaobaoAlitripTravelNormalvisaStoreuserRequest 初始化TaobaoAlitripTravelNormalvisaStoreuserAPIRequest对象
func NewTaobaoAlitripTravelNormalvisaStoreuserRequest() *TaobaoAlitripTravelNormalvisaStoreuserAPIRequest {
	return &TaobaoAlitripTravelNormalvisaStoreuserAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoAlitripTravelNormalvisaStoreuserAPIRequest) Reset() {
	r._normalVisaUserUnitList = r._normalVisaUserUnitList[:0]
	r._bizOrderId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripTravelNormalvisaStoreuserAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.travel.normalvisa.storeuser"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripTravelNormalvisaStoreuserAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAlitripTravelNormalvisaStoreuserAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNormalVisaUserUnitList is NormalVisaUserUnitList Setter
// 列表：签证人信息列表
func (r *TaobaoAlitripTravelNormalvisaStoreuserAPIRequest) SetNormalVisaUserUnitList(_normalVisaUserUnitList []NormalVisaUserUnit) error {
	r._normalVisaUserUnitList = _normalVisaUserUnitList
	r.Set("normal_visa_user_unit_list", _normalVisaUserUnitList)
	return nil
}

// GetNormalVisaUserUnitList NormalVisaUserUnitList Getter
func (r TaobaoAlitripTravelNormalvisaStoreuserAPIRequest) GetNormalVisaUserUnitList() []NormalVisaUserUnit {
	return r._normalVisaUserUnitList
}

// SetBizOrderId is BizOrderId Setter
// 订单id
func (r *TaobaoAlitripTravelNormalvisaStoreuserAPIRequest) SetBizOrderId(_bizOrderId int64) error {
	r._bizOrderId = _bizOrderId
	r.Set("biz_order_id", _bizOrderId)
	return nil
}

// GetBizOrderId BizOrderId Getter
func (r TaobaoAlitripTravelNormalvisaStoreuserAPIRequest) GetBizOrderId() int64 {
	return r._bizOrderId
}

var poolTaobaoAlitripTravelNormalvisaStoreuserAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoAlitripTravelNormalvisaStoreuserRequest()
	},
}

// GetTaobaoAlitripTravelNormalvisaStoreuserRequest 从 sync.Pool 获取 TaobaoAlitripTravelNormalvisaStoreuserAPIRequest
func GetTaobaoAlitripTravelNormalvisaStoreuserAPIRequest() *TaobaoAlitripTravelNormalvisaStoreuserAPIRequest {
	return poolTaobaoAlitripTravelNormalvisaStoreuserAPIRequest.Get().(*TaobaoAlitripTravelNormalvisaStoreuserAPIRequest)
}

// ReleaseTaobaoAlitripTravelNormalvisaStoreuserAPIRequest 将 TaobaoAlitripTravelNormalvisaStoreuserAPIRequest 放入 sync.Pool
func ReleaseTaobaoAlitripTravelNormalvisaStoreuserAPIRequest(v *TaobaoAlitripTravelNormalvisaStoreuserAPIRequest) {
	v.Reset()
	poolTaobaoAlitripTravelNormalvisaStoreuserAPIRequest.Put(v)
}

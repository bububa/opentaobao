package xhotelitem

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelRateplanDeleteAPIRequest 价格计划rateplan删除 API请求
// taobao.xhotel.rateplan.delete
//
// 酒店产品库rateplan删除，同时删除级联的rate，请谨慎使用
type TaobaoXhotelRateplanDeleteAPIRequest struct {
	model.Params
	// 系统商，一般不用填写，使用须申请
	_vendor string
	// 商家价格政策编码
	_rateplanCode string
	// ratepland标识
	_rpId int64
}

// NewTaobaoXhotelRateplanDeleteRequest 初始化TaobaoXhotelRateplanDeleteAPIRequest对象
func NewTaobaoXhotelRateplanDeleteRequest() *TaobaoXhotelRateplanDeleteAPIRequest {
	return &TaobaoXhotelRateplanDeleteAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoXhotelRateplanDeleteAPIRequest) Reset() {
	r._vendor = ""
	r._rateplanCode = ""
	r._rpId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoXhotelRateplanDeleteAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.rateplan.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoXhotelRateplanDeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoXhotelRateplanDeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetVendor is Vendor Setter
// 系统商，一般不用填写，使用须申请
func (r *TaobaoXhotelRateplanDeleteAPIRequest) SetVendor(_vendor string) error {
	r._vendor = _vendor
	r.Set("vendor", _vendor)
	return nil
}

// GetVendor Vendor Getter
func (r TaobaoXhotelRateplanDeleteAPIRequest) GetVendor() string {
	return r._vendor
}

// SetRateplanCode is RateplanCode Setter
// 商家价格政策编码
func (r *TaobaoXhotelRateplanDeleteAPIRequest) SetRateplanCode(_rateplanCode string) error {
	r._rateplanCode = _rateplanCode
	r.Set("rateplan_code", _rateplanCode)
	return nil
}

// GetRateplanCode RateplanCode Getter
func (r TaobaoXhotelRateplanDeleteAPIRequest) GetRateplanCode() string {
	return r._rateplanCode
}

// SetRpId is RpId Setter
// ratepland标识
func (r *TaobaoXhotelRateplanDeleteAPIRequest) SetRpId(_rpId int64) error {
	r._rpId = _rpId
	r.Set("rp_id", _rpId)
	return nil
}

// GetRpId RpId Getter
func (r TaobaoXhotelRateplanDeleteAPIRequest) GetRpId() int64 {
	return r._rpId
}

var poolTaobaoXhotelRateplanDeleteAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoXhotelRateplanDeleteRequest()
	},
}

// GetTaobaoXhotelRateplanDeleteRequest 从 sync.Pool 获取 TaobaoXhotelRateplanDeleteAPIRequest
func GetTaobaoXhotelRateplanDeleteAPIRequest() *TaobaoXhotelRateplanDeleteAPIRequest {
	return poolTaobaoXhotelRateplanDeleteAPIRequest.Get().(*TaobaoXhotelRateplanDeleteAPIRequest)
}

// ReleaseTaobaoXhotelRateplanDeleteAPIRequest 将 TaobaoXhotelRateplanDeleteAPIRequest 放入 sync.Pool
func ReleaseTaobaoXhotelRateplanDeleteAPIRequest(v *TaobaoXhotelRateplanDeleteAPIRequest) {
	v.Reset()
	poolTaobaoXhotelRateplanDeleteAPIRequest.Put(v)
}

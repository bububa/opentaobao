package xhotelitem

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelRateDeleteAPIRequest rate删除接口 API请求
// taobao.xhotel.rate.delete
//
// 酒店产品库rate删除
type TaobaoXhotelRateDeleteAPIRequest struct {
	model.Params
	// 系统商，一般不用填写，使用须申请
	_vendor string
	// 商家价格政策编码
	_rateplanCode string
	// 商家房型ID
	_outRid string
}

// NewTaobaoXhotelRateDeleteRequest 初始化TaobaoXhotelRateDeleteAPIRequest对象
func NewTaobaoXhotelRateDeleteRequest() *TaobaoXhotelRateDeleteAPIRequest {
	return &TaobaoXhotelRateDeleteAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoXhotelRateDeleteAPIRequest) Reset() {
	r._vendor = ""
	r._rateplanCode = ""
	r._outRid = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoXhotelRateDeleteAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.rate.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoXhotelRateDeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoXhotelRateDeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetVendor is Vendor Setter
// 系统商，一般不用填写，使用须申请
func (r *TaobaoXhotelRateDeleteAPIRequest) SetVendor(_vendor string) error {
	r._vendor = _vendor
	r.Set("vendor", _vendor)
	return nil
}

// GetVendor Vendor Getter
func (r TaobaoXhotelRateDeleteAPIRequest) GetVendor() string {
	return r._vendor
}

// SetRateplanCode is RateplanCode Setter
// 商家价格政策编码
func (r *TaobaoXhotelRateDeleteAPIRequest) SetRateplanCode(_rateplanCode string) error {
	r._rateplanCode = _rateplanCode
	r.Set("rateplan_code", _rateplanCode)
	return nil
}

// GetRateplanCode RateplanCode Getter
func (r TaobaoXhotelRateDeleteAPIRequest) GetRateplanCode() string {
	return r._rateplanCode
}

// SetOutRid is OutRid Setter
// 商家房型ID
func (r *TaobaoXhotelRateDeleteAPIRequest) SetOutRid(_outRid string) error {
	r._outRid = _outRid
	r.Set("out_rid", _outRid)
	return nil
}

// GetOutRid OutRid Getter
func (r TaobaoXhotelRateDeleteAPIRequest) GetOutRid() string {
	return r._outRid
}

var poolTaobaoXhotelRateDeleteAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoXhotelRateDeleteRequest()
	},
}

// GetTaobaoXhotelRateDeleteRequest 从 sync.Pool 获取 TaobaoXhotelRateDeleteAPIRequest
func GetTaobaoXhotelRateDeleteAPIRequest() *TaobaoXhotelRateDeleteAPIRequest {
	return poolTaobaoXhotelRateDeleteAPIRequest.Get().(*TaobaoXhotelRateDeleteAPIRequest)
}

// ReleaseTaobaoXhotelRateDeleteAPIRequest 将 TaobaoXhotelRateDeleteAPIRequest 放入 sync.Pool
func ReleaseTaobaoXhotelRateDeleteAPIRequest(v *TaobaoXhotelRateDeleteAPIRequest) {
	v.Reset()
	poolTaobaoXhotelRateDeleteAPIRequest.Put(v)
}

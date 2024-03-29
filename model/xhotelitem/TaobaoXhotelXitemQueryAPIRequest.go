package xhotelitem

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelXitemQueryAPIRequest 查询 x 元素 API请求
// taobao.xhotel.xitem.query
//
// 查询 x 元素
type TaobaoXhotelXitemQueryAPIRequest struct {
	model.Params
	// 系统商，一般不填写，使用须申请
	_vendor string
	// 商家酒店ID，指明该 x 元素属于哪家酒店
	_outHid string
	// 商家房型ID，指明该 x 元素关联哪个房型
	_outRid string
	// 商家 RP ID，指明该 x 元素关联了哪个 RP
	_ratePlanCode string
	// 需要查询的 x_code 编码
	_outXCodes string
}

// NewTaobaoXhotelXitemQueryRequest 初始化TaobaoXhotelXitemQueryAPIRequest对象
func NewTaobaoXhotelXitemQueryRequest() *TaobaoXhotelXitemQueryAPIRequest {
	return &TaobaoXhotelXitemQueryAPIRequest{
		Params: model.NewParams(5),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoXhotelXitemQueryAPIRequest) Reset() {
	r._vendor = ""
	r._outHid = ""
	r._outRid = ""
	r._ratePlanCode = ""
	r._outXCodes = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoXhotelXitemQueryAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.xitem.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoXhotelXitemQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoXhotelXitemQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetVendor is Vendor Setter
// 系统商，一般不填写，使用须申请
func (r *TaobaoXhotelXitemQueryAPIRequest) SetVendor(_vendor string) error {
	r._vendor = _vendor
	r.Set("vendor", _vendor)
	return nil
}

// GetVendor Vendor Getter
func (r TaobaoXhotelXitemQueryAPIRequest) GetVendor() string {
	return r._vendor
}

// SetOutHid is OutHid Setter
// 商家酒店ID，指明该 x 元素属于哪家酒店
func (r *TaobaoXhotelXitemQueryAPIRequest) SetOutHid(_outHid string) error {
	r._outHid = _outHid
	r.Set("out_hid", _outHid)
	return nil
}

// GetOutHid OutHid Getter
func (r TaobaoXhotelXitemQueryAPIRequest) GetOutHid() string {
	return r._outHid
}

// SetOutRid is OutRid Setter
// 商家房型ID，指明该 x 元素关联哪个房型
func (r *TaobaoXhotelXitemQueryAPIRequest) SetOutRid(_outRid string) error {
	r._outRid = _outRid
	r.Set("out_rid", _outRid)
	return nil
}

// GetOutRid OutRid Getter
func (r TaobaoXhotelXitemQueryAPIRequest) GetOutRid() string {
	return r._outRid
}

// SetRatePlanCode is RatePlanCode Setter
// 商家 RP ID，指明该 x 元素关联了哪个 RP
func (r *TaobaoXhotelXitemQueryAPIRequest) SetRatePlanCode(_ratePlanCode string) error {
	r._ratePlanCode = _ratePlanCode
	r.Set("rate_plan_code", _ratePlanCode)
	return nil
}

// GetRatePlanCode RatePlanCode Getter
func (r TaobaoXhotelXitemQueryAPIRequest) GetRatePlanCode() string {
	return r._ratePlanCode
}

// SetOutXCodes is OutXCodes Setter
// 需要查询的 x_code 编码
func (r *TaobaoXhotelXitemQueryAPIRequest) SetOutXCodes(_outXCodes string) error {
	r._outXCodes = _outXCodes
	r.Set("out_x_codes", _outXCodes)
	return nil
}

// GetOutXCodes OutXCodes Getter
func (r TaobaoXhotelXitemQueryAPIRequest) GetOutXCodes() string {
	return r._outXCodes
}

var poolTaobaoXhotelXitemQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoXhotelXitemQueryRequest()
	},
}

// GetTaobaoXhotelXitemQueryRequest 从 sync.Pool 获取 TaobaoXhotelXitemQueryAPIRequest
func GetTaobaoXhotelXitemQueryAPIRequest() *TaobaoXhotelXitemQueryAPIRequest {
	return poolTaobaoXhotelXitemQueryAPIRequest.Get().(*TaobaoXhotelXitemQueryAPIRequest)
}

// ReleaseTaobaoXhotelXitemQueryAPIRequest 将 TaobaoXhotelXitemQueryAPIRequest 放入 sync.Pool
func ReleaseTaobaoXhotelXitemQueryAPIRequest(v *TaobaoXhotelXitemQueryAPIRequest) {
	v.Reset()
	poolTaobaoXhotelXitemQueryAPIRequest.Put(v)
}

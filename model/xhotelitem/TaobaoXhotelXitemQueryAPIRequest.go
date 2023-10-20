package xhotelitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoxhotelxitemqueryAPIRequest 查询 x 元素 API请求
// taobao.xhotel.xitem.query
//
// 查询 x 元素
type TaobaoxhotelxitemqueryAPIRequest struct {
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

// NewTaobaoxhotelxitemqueryRequest 初始化TaobaoxhotelxitemqueryAPIRequest对象
func NewTaobaoxhotelxitemqueryRequest() *TaobaoxhotelxitemqueryAPIRequest {
	return &TaobaoxhotelxitemqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoxhotelxitemqueryAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.xitem.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoxhotelxitemqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoxhotelxitemqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetVendor is Vendor Setter
// 系统商，一般不填写，使用须申请
func (r *TaobaoxhotelxitemqueryAPIRequest) SetVendor(_vendor string) error {
	r._vendor = _vendor
	r.Set("vendor", _vendor)
	return nil
}

// GetVendor Vendor Getter
func (r TaobaoxhotelxitemqueryAPIRequest) GetVendor() string {
	return r._vendor
}

// SetOutHid is OutHid Setter
// 商家酒店ID，指明该 x 元素属于哪家酒店
func (r *TaobaoxhotelxitemqueryAPIRequest) SetOutHid(_outHid string) error {
	r._outHid = _outHid
	r.Set("out_hid", _outHid)
	return nil
}

// GetOutHid OutHid Getter
func (r TaobaoxhotelxitemqueryAPIRequest) GetOutHid() string {
	return r._outHid
}

// SetOutRid is OutRid Setter
// 商家房型ID，指明该 x 元素关联哪个房型
func (r *TaobaoxhotelxitemqueryAPIRequest) SetOutRid(_outRid string) error {
	r._outRid = _outRid
	r.Set("out_rid", _outRid)
	return nil
}

// GetOutRid OutRid Getter
func (r TaobaoxhotelxitemqueryAPIRequest) GetOutRid() string {
	return r._outRid
}

// SetRatePlanCode is RatePlanCode Setter
// 商家 RP ID，指明该 x 元素关联了哪个 RP
func (r *TaobaoxhotelxitemqueryAPIRequest) SetRatePlanCode(_ratePlanCode string) error {
	r._ratePlanCode = _ratePlanCode
	r.Set("rate_plan_code", _ratePlanCode)
	return nil
}

// GetRatePlanCode RatePlanCode Getter
func (r TaobaoxhotelxitemqueryAPIRequest) GetRatePlanCode() string {
	return r._ratePlanCode
}

// SetOutXCodes is OutXCodes Setter
// 需要查询的 x_code 编码
func (r *TaobaoxhotelxitemqueryAPIRequest) SetOutXCodes(_outXCodes string) error {
	r._outXCodes = _outXCodes
	r.Set("out_x_codes", _outXCodes)
	return nil
}

// GetOutXCodes OutXCodes Getter
func (r TaobaoxhotelxitemqueryAPIRequest) GetOutXCodes() string {
	return r._outXCodes
}

package xhotelitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoxhotelxitemdeleteAPIRequest 删除 x 元素 API请求
// taobao.xhotel.xitem.delete
//
// 删除 x 元素
type TaobaoxhotelxitemdeleteAPIRequest struct {
	model.Params
	// 系统商，一般不填写，使用须申请
	_vendor string
	// 商家酒店ID，指明该 x 元素属于哪家酒店
	_outHid string
	// 需要删除的 x_code 编码
	_outXCodes string
}

// NewTaobaoxhotelxitemdeleteRequest 初始化TaobaoxhotelxitemdeleteAPIRequest对象
func NewTaobaoxhotelxitemdeleteRequest() *TaobaoxhotelxitemdeleteAPIRequest {
	return &TaobaoxhotelxitemdeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoxhotelxitemdeleteAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.xitem.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoxhotelxitemdeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoxhotelxitemdeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetVendor is Vendor Setter
// 系统商，一般不填写，使用须申请
func (r *TaobaoxhotelxitemdeleteAPIRequest) SetVendor(_vendor string) error {
	r._vendor = _vendor
	r.Set("vendor", _vendor)
	return nil
}

// GetVendor Vendor Getter
func (r TaobaoxhotelxitemdeleteAPIRequest) GetVendor() string {
	return r._vendor
}

// SetOutHid is OutHid Setter
// 商家酒店ID，指明该 x 元素属于哪家酒店
func (r *TaobaoxhotelxitemdeleteAPIRequest) SetOutHid(_outHid string) error {
	r._outHid = _outHid
	r.Set("out_hid", _outHid)
	return nil
}

// GetOutHid OutHid Getter
func (r TaobaoxhotelxitemdeleteAPIRequest) GetOutHid() string {
	return r._outHid
}

// SetOutXCodes is OutXCodes Setter
// 需要删除的 x_code 编码
func (r *TaobaoxhotelxitemdeleteAPIRequest) SetOutXCodes(_outXCodes string) error {
	r._outXCodes = _outXCodes
	r.Set("out_x_codes", _outXCodes)
	return nil
}

// GetOutXCodes OutXCodes Getter
func (r TaobaoxhotelxitemdeleteAPIRequest) GetOutXCodes() string {
	return r._outXCodes
}

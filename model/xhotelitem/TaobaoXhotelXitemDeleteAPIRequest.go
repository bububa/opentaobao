package xhotelitem

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelXitemDeleteAPIRequest 删除 x 元素 API请求
// taobao.xhotel.xitem.delete
//
// 删除 x 元素
type TaobaoXhotelXitemDeleteAPIRequest struct {
	model.Params
	// 系统商，一般不填写，使用须申请
	_vendor string
	// 商家酒店ID，指明该 x 元素属于哪家酒店
	_outHid string
	// 需要删除的 x_code 编码
	_outXCodes string
}

// NewTaobaoXhotelXitemDeleteRequest 初始化TaobaoXhotelXitemDeleteAPIRequest对象
func NewTaobaoXhotelXitemDeleteRequest() *TaobaoXhotelXitemDeleteAPIRequest {
	return &TaobaoXhotelXitemDeleteAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoXhotelXitemDeleteAPIRequest) Reset() {
	r._vendor = ""
	r._outHid = ""
	r._outXCodes = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoXhotelXitemDeleteAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.xitem.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoXhotelXitemDeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoXhotelXitemDeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetVendor is Vendor Setter
// 系统商，一般不填写，使用须申请
func (r *TaobaoXhotelXitemDeleteAPIRequest) SetVendor(_vendor string) error {
	r._vendor = _vendor
	r.Set("vendor", _vendor)
	return nil
}

// GetVendor Vendor Getter
func (r TaobaoXhotelXitemDeleteAPIRequest) GetVendor() string {
	return r._vendor
}

// SetOutHid is OutHid Setter
// 商家酒店ID，指明该 x 元素属于哪家酒店
func (r *TaobaoXhotelXitemDeleteAPIRequest) SetOutHid(_outHid string) error {
	r._outHid = _outHid
	r.Set("out_hid", _outHid)
	return nil
}

// GetOutHid OutHid Getter
func (r TaobaoXhotelXitemDeleteAPIRequest) GetOutHid() string {
	return r._outHid
}

// SetOutXCodes is OutXCodes Setter
// 需要删除的 x_code 编码
func (r *TaobaoXhotelXitemDeleteAPIRequest) SetOutXCodes(_outXCodes string) error {
	r._outXCodes = _outXCodes
	r.Set("out_x_codes", _outXCodes)
	return nil
}

// GetOutXCodes OutXCodes Getter
func (r TaobaoXhotelXitemDeleteAPIRequest) GetOutXCodes() string {
	return r._outXCodes
}

var poolTaobaoXhotelXitemDeleteAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoXhotelXitemDeleteRequest()
	},
}

// GetTaobaoXhotelXitemDeleteRequest 从 sync.Pool 获取 TaobaoXhotelXitemDeleteAPIRequest
func GetTaobaoXhotelXitemDeleteAPIRequest() *TaobaoXhotelXitemDeleteAPIRequest {
	return poolTaobaoXhotelXitemDeleteAPIRequest.Get().(*TaobaoXhotelXitemDeleteAPIRequest)
}

// ReleaseTaobaoXhotelXitemDeleteAPIRequest 将 TaobaoXhotelXitemDeleteAPIRequest 放入 sync.Pool
func ReleaseTaobaoXhotelXitemDeleteAPIRequest(v *TaobaoXhotelXitemDeleteAPIRequest) {
	v.Reset()
	poolTaobaoXhotelXitemDeleteAPIRequest.Put(v)
}

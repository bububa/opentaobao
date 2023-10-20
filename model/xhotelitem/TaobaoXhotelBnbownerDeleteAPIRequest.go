package xhotelitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoxhotelbnbownerdeleteAPIRequest 民宿房东删除接口 API请求
// taobao.xhotel.bnbowner.delete
//
// 民宿房东删除接口，删除房东后，对应的门店及房源会同步删除，请谨慎使用
type TaobaoxhotelbnbownerdeleteAPIRequest struct {
	model.Params
	// 对接系统商名称：可为空不要乱填，需要申请后使用，默认taobao
	_vendor string
	// 房东Id，系统商outer_id
	_outerId string
}

// NewTaobaoxhotelbnbownerdeleteRequest 初始化TaobaoxhotelbnbownerdeleteAPIRequest对象
func NewTaobaoxhotelbnbownerdeleteRequest() *TaobaoxhotelbnbownerdeleteAPIRequest {
	return &TaobaoxhotelbnbownerdeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoxhotelbnbownerdeleteAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.bnbowner.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoxhotelbnbownerdeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoxhotelbnbownerdeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetVendor is Vendor Setter
// 对接系统商名称：可为空不要乱填，需要申请后使用，默认taobao
func (r *TaobaoxhotelbnbownerdeleteAPIRequest) SetVendor(_vendor string) error {
	r._vendor = _vendor
	r.Set("vendor", _vendor)
	return nil
}

// GetVendor Vendor Getter
func (r TaobaoxhotelbnbownerdeleteAPIRequest) GetVendor() string {
	return r._vendor
}

// SetOuterId is OuterId Setter
// 房东Id，系统商outer_id
func (r *TaobaoxhotelbnbownerdeleteAPIRequest) SetOuterId(_outerId string) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// GetOuterId OuterId Getter
func (r TaobaoxhotelbnbownerdeleteAPIRequest) GetOuterId() string {
	return r._outerId
}

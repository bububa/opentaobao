package xhotelitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoXhotelServicetimeGetAPIRequest
查询实体对应的服务时间数据 API请求
taobao.xhotel.servicetime.get

通过实体来获取对应的服务时间数据 */
type TaobaoXhotelServicetimeGetAPIRequest struct {
	model.Params
	// 酒店id
	_hid int64
}

// NewTaobaoXhotelServicetimeGetRequest 初始化TaobaoXhotelServicetimeGetAPIRequest对象
func NewTaobaoXhotelServicetimeGetRequest() *TaobaoXhotelServicetimeGetAPIRequest {
	return &TaobaoXhotelServicetimeGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoXhotelServicetimeGetAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.servicetime.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoXhotelServicetimeGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Hid Setter
// 酒店id
func (r *TaobaoXhotelServicetimeGetAPIRequest) SetHid(_hid int64) error {
	r._hid = _hid
	r.Set("hid", _hid)
	return nil
}

// Get Hid Getter
func (r TaobaoXhotelServicetimeGetAPIRequest) GetHid() int64 {
	return r._hid
}

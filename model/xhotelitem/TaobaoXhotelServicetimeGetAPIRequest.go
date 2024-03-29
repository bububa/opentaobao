package xhotelitem

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelServicetimeGetAPIRequest 查询实体对应的服务时间数据 API请求
// taobao.xhotel.servicetime.get
//
// 通过实体来获取对应的服务时间数据
type TaobaoXhotelServicetimeGetAPIRequest struct {
	model.Params
	// 酒店id
	_hid int64
}

// NewTaobaoXhotelServicetimeGetRequest 初始化TaobaoXhotelServicetimeGetAPIRequest对象
func NewTaobaoXhotelServicetimeGetRequest() *TaobaoXhotelServicetimeGetAPIRequest {
	return &TaobaoXhotelServicetimeGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoXhotelServicetimeGetAPIRequest) Reset() {
	r._hid = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoXhotelServicetimeGetAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.servicetime.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoXhotelServicetimeGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoXhotelServicetimeGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetHid is Hid Setter
// 酒店id
func (r *TaobaoXhotelServicetimeGetAPIRequest) SetHid(_hid int64) error {
	r._hid = _hid
	r.Set("hid", _hid)
	return nil
}

// GetHid Hid Getter
func (r TaobaoXhotelServicetimeGetAPIRequest) GetHid() int64 {
	return r._hid
}

var poolTaobaoXhotelServicetimeGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoXhotelServicetimeGetRequest()
	},
}

// GetTaobaoXhotelServicetimeGetRequest 从 sync.Pool 获取 TaobaoXhotelServicetimeGetAPIRequest
func GetTaobaoXhotelServicetimeGetAPIRequest() *TaobaoXhotelServicetimeGetAPIRequest {
	return poolTaobaoXhotelServicetimeGetAPIRequest.Get().(*TaobaoXhotelServicetimeGetAPIRequest)
}

// ReleaseTaobaoXhotelServicetimeGetAPIRequest 将 TaobaoXhotelServicetimeGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoXhotelServicetimeGetAPIRequest(v *TaobaoXhotelServicetimeGetAPIRequest) {
	v.Reset()
	poolTaobaoXhotelServicetimeGetAPIRequest.Put(v)
}

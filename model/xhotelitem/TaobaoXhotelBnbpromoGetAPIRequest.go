package xhotelitem

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelBnbpromoGetAPIRequest 民宿查询营销活动 API请求
// taobao.xhotel.bnbpromo.get
//
// 民宿查询营销活动
type TaobaoXhotelBnbpromoGetAPIRequest struct {
	model.Params
	// 外部活动code
	_outerActivityCode string
}

// NewTaobaoXhotelBnbpromoGetRequest 初始化TaobaoXhotelBnbpromoGetAPIRequest对象
func NewTaobaoXhotelBnbpromoGetRequest() *TaobaoXhotelBnbpromoGetAPIRequest {
	return &TaobaoXhotelBnbpromoGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoXhotelBnbpromoGetAPIRequest) Reset() {
	r._outerActivityCode = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoXhotelBnbpromoGetAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.bnbpromo.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoXhotelBnbpromoGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoXhotelBnbpromoGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOuterActivityCode is OuterActivityCode Setter
// 外部活动code
func (r *TaobaoXhotelBnbpromoGetAPIRequest) SetOuterActivityCode(_outerActivityCode string) error {
	r._outerActivityCode = _outerActivityCode
	r.Set("outer_activity_code", _outerActivityCode)
	return nil
}

// GetOuterActivityCode OuterActivityCode Getter
func (r TaobaoXhotelBnbpromoGetAPIRequest) GetOuterActivityCode() string {
	return r._outerActivityCode
}

var poolTaobaoXhotelBnbpromoGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoXhotelBnbpromoGetRequest()
	},
}

// GetTaobaoXhotelBnbpromoGetRequest 从 sync.Pool 获取 TaobaoXhotelBnbpromoGetAPIRequest
func GetTaobaoXhotelBnbpromoGetAPIRequest() *TaobaoXhotelBnbpromoGetAPIRequest {
	return poolTaobaoXhotelBnbpromoGetAPIRequest.Get().(*TaobaoXhotelBnbpromoGetAPIRequest)
}

// ReleaseTaobaoXhotelBnbpromoGetAPIRequest 将 TaobaoXhotelBnbpromoGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoXhotelBnbpromoGetAPIRequest(v *TaobaoXhotelBnbpromoGetAPIRequest) {
	v.Reset()
	poolTaobaoXhotelBnbpromoGetAPIRequest.Put(v)
}

package xhotelonlineorder

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelOrderDetailSearchAPIRequest 订单详情查询 API请求
// taobao.xhotel.order.detail.search
//
// 提供订单详情查询
type TaobaoXhotelOrderDetailSearchAPIRequest struct {
	model.Params
	// 外部订单号
	_outOid string
	// 外部订单号
	_tid int64
}

// NewTaobaoXhotelOrderDetailSearchRequest 初始化TaobaoXhotelOrderDetailSearchAPIRequest对象
func NewTaobaoXhotelOrderDetailSearchRequest() *TaobaoXhotelOrderDetailSearchAPIRequest {
	return &TaobaoXhotelOrderDetailSearchAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoXhotelOrderDetailSearchAPIRequest) Reset() {
	r._outOid = ""
	r._tid = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoXhotelOrderDetailSearchAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.order.detail.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoXhotelOrderDetailSearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoXhotelOrderDetailSearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOutOid is OutOid Setter
// 外部订单号
func (r *TaobaoXhotelOrderDetailSearchAPIRequest) SetOutOid(_outOid string) error {
	r._outOid = _outOid
	r.Set("out_oid", _outOid)
	return nil
}

// GetOutOid OutOid Getter
func (r TaobaoXhotelOrderDetailSearchAPIRequest) GetOutOid() string {
	return r._outOid
}

// SetTid is Tid Setter
// 外部订单号
func (r *TaobaoXhotelOrderDetailSearchAPIRequest) SetTid(_tid int64) error {
	r._tid = _tid
	r.Set("tid", _tid)
	return nil
}

// GetTid Tid Getter
func (r TaobaoXhotelOrderDetailSearchAPIRequest) GetTid() int64 {
	return r._tid
}

var poolTaobaoXhotelOrderDetailSearchAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoXhotelOrderDetailSearchRequest()
	},
}

// GetTaobaoXhotelOrderDetailSearchRequest 从 sync.Pool 获取 TaobaoXhotelOrderDetailSearchAPIRequest
func GetTaobaoXhotelOrderDetailSearchAPIRequest() *TaobaoXhotelOrderDetailSearchAPIRequest {
	return poolTaobaoXhotelOrderDetailSearchAPIRequest.Get().(*TaobaoXhotelOrderDetailSearchAPIRequest)
}

// ReleaseTaobaoXhotelOrderDetailSearchAPIRequest 将 TaobaoXhotelOrderDetailSearchAPIRequest 放入 sync.Pool
func ReleaseTaobaoXhotelOrderDetailSearchAPIRequest(v *TaobaoXhotelOrderDetailSearchAPIRequest) {
	v.Reset()
	poolTaobaoXhotelOrderDetailSearchAPIRequest.Put(v)
}

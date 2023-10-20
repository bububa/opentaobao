package tmallnr

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallNrOrderQueryJstAPIRequest 获取同城配送业务单笔订单 API请求
// tmall.nr.order.query.jst
//
// 同城配送业务获取单笔订单
type TmallNrOrderQueryJstAPIRequest struct {
	model.Params
	// 业务标识，dss标识定时送业务；jsd表示极速达业务
	_bizIdentity string
	// 预留-扩展信息
	_extParam string
	// 交易主订单号
	_orderId int64
}

// NewTmallNrOrderQueryJstRequest 初始化TmallNrOrderQueryJstAPIRequest对象
func NewTmallNrOrderQueryJstRequest() *TmallNrOrderQueryJstAPIRequest {
	return &TmallNrOrderQueryJstAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallNrOrderQueryJstAPIRequest) Reset() {
	r._bizIdentity = ""
	r._extParam = ""
	r._orderId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallNrOrderQueryJstAPIRequest) GetApiMethodName() string {
	return "tmall.nr.order.query.jst"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallNrOrderQueryJstAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallNrOrderQueryJstAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBizIdentity is BizIdentity Setter
// 业务标识，dss标识定时送业务；jsd表示极速达业务
func (r *TmallNrOrderQueryJstAPIRequest) SetBizIdentity(_bizIdentity string) error {
	r._bizIdentity = _bizIdentity
	r.Set("biz_identity", _bizIdentity)
	return nil
}

// GetBizIdentity BizIdentity Getter
func (r TmallNrOrderQueryJstAPIRequest) GetBizIdentity() string {
	return r._bizIdentity
}

// SetExtParam is ExtParam Setter
// 预留-扩展信息
func (r *TmallNrOrderQueryJstAPIRequest) SetExtParam(_extParam string) error {
	r._extParam = _extParam
	r.Set("ext_param", _extParam)
	return nil
}

// GetExtParam ExtParam Getter
func (r TmallNrOrderQueryJstAPIRequest) GetExtParam() string {
	return r._extParam
}

// SetOrderId is OrderId Setter
// 交易主订单号
func (r *TmallNrOrderQueryJstAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r TmallNrOrderQueryJstAPIRequest) GetOrderId() int64 {
	return r._orderId
}

var poolTmallNrOrderQueryJstAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallNrOrderQueryJstRequest()
	},
}

// GetTmallNrOrderQueryJstRequest 从 sync.Pool 获取 TmallNrOrderQueryJstAPIRequest
func GetTmallNrOrderQueryJstAPIRequest() *TmallNrOrderQueryJstAPIRequest {
	return poolTmallNrOrderQueryJstAPIRequest.Get().(*TmallNrOrderQueryJstAPIRequest)
}

// ReleaseTmallNrOrderQueryJstAPIRequest 将 TmallNrOrderQueryJstAPIRequest 放入 sync.Pool
func ReleaseTmallNrOrderQueryJstAPIRequest(v *TmallNrOrderQueryJstAPIRequest) {
	v.Reset()
	poolTmallNrOrderQueryJstAPIRequest.Put(v)
}

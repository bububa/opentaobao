package tmallnr

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallnrorderqueryjstAPIRequest 获取同城配送业务单笔订单 API请求
// tmall.nr.order.query.jst
//
// 同城配送业务获取单笔订单
type TmallnrorderqueryjstAPIRequest struct {
	model.Params
	// 业务标识，dss标识定时送业务；jsd表示极速达业务
	_bizIdentity string
	// 预留-扩展信息
	_extParam string
	// 交易主订单号
	_orderId int64
}

// NewTmallnrorderqueryjstRequest 初始化TmallnrorderqueryjstAPIRequest对象
func NewTmallnrorderqueryjstRequest() *TmallnrorderqueryjstAPIRequest {
	return &TmallnrorderqueryjstAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallnrorderqueryjstAPIRequest) GetApiMethodName() string {
	return "tmall.nr.order.query.jst"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallnrorderqueryjstAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallnrorderqueryjstAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBizIdentity is BizIdentity Setter
// 业务标识，dss标识定时送业务；jsd表示极速达业务
func (r *TmallnrorderqueryjstAPIRequest) SetBizIdentity(_bizIdentity string) error {
	r._bizIdentity = _bizIdentity
	r.Set("biz_identity", _bizIdentity)
	return nil
}

// GetBizIdentity BizIdentity Getter
func (r TmallnrorderqueryjstAPIRequest) GetBizIdentity() string {
	return r._bizIdentity
}

// SetExtParam is ExtParam Setter
// 预留-扩展信息
func (r *TmallnrorderqueryjstAPIRequest) SetExtParam(_extParam string) error {
	r._extParam = _extParam
	r.Set("ext_param", _extParam)
	return nil
}

// GetExtParam ExtParam Getter
func (r TmallnrorderqueryjstAPIRequest) GetExtParam() string {
	return r._extParam
}

// SetOrderId is OrderId Setter
// 交易主订单号
func (r *TmallnrorderqueryjstAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r TmallnrorderqueryjstAPIRequest) GetOrderId() int64 {
	return r._orderId
}

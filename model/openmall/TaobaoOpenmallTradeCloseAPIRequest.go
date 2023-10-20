package openmall

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoopenmalltradecloseAPIRequest 关闭订单 API请求
// taobao.openmall.trade.close
//
// 关闭订单
type TaobaoopenmalltradecloseAPIRequest struct {
	model.Params
	// 分销者信息
	_distributor string
	// 关单原因
	_reason string
	// 淘宝订单号
	_tid int64
}

// NewTaobaoopenmalltradecloseRequest 初始化TaobaoopenmalltradecloseAPIRequest对象
func NewTaobaoopenmalltradecloseRequest() *TaobaoopenmalltradecloseAPIRequest {
	return &TaobaoopenmalltradecloseAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoopenmalltradecloseAPIRequest) GetApiMethodName() string {
	return "taobao.openmall.trade.close"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoopenmalltradecloseAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoopenmalltradecloseAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDistributor is Distributor Setter
// 分销者信息
func (r *TaobaoopenmalltradecloseAPIRequest) SetDistributor(_distributor string) error {
	r._distributor = _distributor
	r.Set("distributor", _distributor)
	return nil
}

// GetDistributor Distributor Getter
func (r TaobaoopenmalltradecloseAPIRequest) GetDistributor() string {
	return r._distributor
}

// SetReason is Reason Setter
// 关单原因
func (r *TaobaoopenmalltradecloseAPIRequest) SetReason(_reason string) error {
	r._reason = _reason
	r.Set("reason", _reason)
	return nil
}

// GetReason Reason Getter
func (r TaobaoopenmalltradecloseAPIRequest) GetReason() string {
	return r._reason
}

// SetTid is Tid Setter
// 淘宝订单号
func (r *TaobaoopenmalltradecloseAPIRequest) SetTid(_tid int64) error {
	r._tid = _tid
	r.Set("tid", _tid)
	return nil
}

// GetTid Tid Getter
func (r TaobaoopenmalltradecloseAPIRequest) GetTid() int64 {
	return r._tid
}

package openmall

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoopenmalltracesearchAPIRequest 获取Openmall订单物流流转信息 API请求
// taobao.openmall.trace.search
//
// 获取Openmall订单物流流转信息
type TaobaoopenmalltracesearchAPIRequest struct {
	model.Params
	// 签约支付宝代扣的账号
	_distributor string
	// 淘宝订单编号
	_tid int64
}

// NewTaobaoopenmalltracesearchRequest 初始化TaobaoopenmalltracesearchAPIRequest对象
func NewTaobaoopenmalltracesearchRequest() *TaobaoopenmalltracesearchAPIRequest {
	return &TaobaoopenmalltracesearchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoopenmalltracesearchAPIRequest) GetApiMethodName() string {
	return "taobao.openmall.trace.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoopenmalltracesearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoopenmalltracesearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDistributor is Distributor Setter
// 签约支付宝代扣的账号
func (r *TaobaoopenmalltracesearchAPIRequest) SetDistributor(_distributor string) error {
	r._distributor = _distributor
	r.Set("distributor", _distributor)
	return nil
}

// GetDistributor Distributor Getter
func (r TaobaoopenmalltracesearchAPIRequest) GetDistributor() string {
	return r._distributor
}

// SetTid is Tid Setter
// 淘宝订单编号
func (r *TaobaoopenmalltracesearchAPIRequest) SetTid(_tid int64) error {
	r._tid = _tid
	r.Set("tid", _tid)
	return nil
}

// GetTid Tid Getter
func (r TaobaoopenmalltracesearchAPIRequest) GetTid() int64 {
	return r._tid
}

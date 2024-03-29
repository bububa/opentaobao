package baoxian

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlipayBaoxianClaimReturngoodsstatusUpdateAPIRequest 更新理赔单退货货物状态 API请求
// alipay.baoxian.claim.returngoodsstatus.update
//
// 更新理赔单退货货物状态
type AlipayBaoxianClaimReturngoodsstatusUpdateAPIRequest struct {
	model.Params
	// 理赔单号
	_claimNo string
	// 退货货物状态
	_goodsStatus string
}

// NewAlipayBaoxianClaimReturngoodsstatusUpdateRequest 初始化AlipayBaoxianClaimReturngoodsstatusUpdateAPIRequest对象
func NewAlipayBaoxianClaimReturngoodsstatusUpdateRequest() *AlipayBaoxianClaimReturngoodsstatusUpdateAPIRequest {
	return &AlipayBaoxianClaimReturngoodsstatusUpdateAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlipayBaoxianClaimReturngoodsstatusUpdateAPIRequest) Reset() {
	r._claimNo = ""
	r._goodsStatus = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlipayBaoxianClaimReturngoodsstatusUpdateAPIRequest) GetApiMethodName() string {
	return "alipay.baoxian.claim.returngoodsstatus.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlipayBaoxianClaimReturngoodsstatusUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlipayBaoxianClaimReturngoodsstatusUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetClaimNo is ClaimNo Setter
// 理赔单号
func (r *AlipayBaoxianClaimReturngoodsstatusUpdateAPIRequest) SetClaimNo(_claimNo string) error {
	r._claimNo = _claimNo
	r.Set("claim_no", _claimNo)
	return nil
}

// GetClaimNo ClaimNo Getter
func (r AlipayBaoxianClaimReturngoodsstatusUpdateAPIRequest) GetClaimNo() string {
	return r._claimNo
}

// SetGoodsStatus is GoodsStatus Setter
// 退货货物状态
func (r *AlipayBaoxianClaimReturngoodsstatusUpdateAPIRequest) SetGoodsStatus(_goodsStatus string) error {
	r._goodsStatus = _goodsStatus
	r.Set("goods_status", _goodsStatus)
	return nil
}

// GetGoodsStatus GoodsStatus Getter
func (r AlipayBaoxianClaimReturngoodsstatusUpdateAPIRequest) GetGoodsStatus() string {
	return r._goodsStatus
}

var poolAlipayBaoxianClaimReturngoodsstatusUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlipayBaoxianClaimReturngoodsstatusUpdateRequest()
	},
}

// GetAlipayBaoxianClaimReturngoodsstatusUpdateRequest 从 sync.Pool 获取 AlipayBaoxianClaimReturngoodsstatusUpdateAPIRequest
func GetAlipayBaoxianClaimReturngoodsstatusUpdateAPIRequest() *AlipayBaoxianClaimReturngoodsstatusUpdateAPIRequest {
	return poolAlipayBaoxianClaimReturngoodsstatusUpdateAPIRequest.Get().(*AlipayBaoxianClaimReturngoodsstatusUpdateAPIRequest)
}

// ReleaseAlipayBaoxianClaimReturngoodsstatusUpdateAPIRequest 将 AlipayBaoxianClaimReturngoodsstatusUpdateAPIRequest 放入 sync.Pool
func ReleaseAlipayBaoxianClaimReturngoodsstatusUpdateAPIRequest(v *AlipayBaoxianClaimReturngoodsstatusUpdateAPIRequest) {
	v.Reset()
	poolAlipayBaoxianClaimReturngoodsstatusUpdateAPIRequest.Put(v)
}

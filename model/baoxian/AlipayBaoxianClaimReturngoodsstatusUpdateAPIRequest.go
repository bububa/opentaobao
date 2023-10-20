package baoxian

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlipaybaoxianclaimreturngoodsstatusupdateAPIRequest 更新理赔单退货货物状态 API请求
// alipay.baoxian.claim.returngoodsstatus.update
//
// 更新理赔单退货货物状态
type AlipaybaoxianclaimreturngoodsstatusupdateAPIRequest struct {
	model.Params
	// 理赔单号
	_claimNo string
	// 退货货物状态
	_goodsStatus string
}

// NewAlipaybaoxianclaimreturngoodsstatusupdateRequest 初始化AlipaybaoxianclaimreturngoodsstatusupdateAPIRequest对象
func NewAlipaybaoxianclaimreturngoodsstatusupdateRequest() *AlipaybaoxianclaimreturngoodsstatusupdateAPIRequest {
	return &AlipaybaoxianclaimreturngoodsstatusupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlipaybaoxianclaimreturngoodsstatusupdateAPIRequest) GetApiMethodName() string {
	return "alipay.baoxian.claim.returngoodsstatus.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlipaybaoxianclaimreturngoodsstatusupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlipaybaoxianclaimreturngoodsstatusupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetClaimNo is ClaimNo Setter
// 理赔单号
func (r *AlipaybaoxianclaimreturngoodsstatusupdateAPIRequest) SetClaimNo(_claimNo string) error {
	r._claimNo = _claimNo
	r.Set("claim_no", _claimNo)
	return nil
}

// GetClaimNo ClaimNo Getter
func (r AlipaybaoxianclaimreturngoodsstatusupdateAPIRequest) GetClaimNo() string {
	return r._claimNo
}

// SetGoodsStatus is GoodsStatus Setter
// 退货货物状态
func (r *AlipaybaoxianclaimreturngoodsstatusupdateAPIRequest) SetGoodsStatus(_goodsStatus string) error {
	r._goodsStatus = _goodsStatus
	r.Set("goods_status", _goodsStatus)
	return nil
}

// GetGoodsStatus GoodsStatus Getter
func (r AlipaybaoxianclaimreturngoodsstatusupdateAPIRequest) GetGoodsStatus() string {
	return r._goodsStatus
}

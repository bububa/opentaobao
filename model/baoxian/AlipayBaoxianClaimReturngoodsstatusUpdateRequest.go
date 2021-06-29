package baoxian

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
更新理赔单退货货物状态 API请求
alipay.baoxian.claim.returngoodsstatus.update

更新理赔单退货货物状态
*/
type AlipayBaoxianClaimReturngoodsstatusUpdateRequest struct {
    model.Params
    // 理赔单号
    _claimNo   string
    // 退货货物状态
    _goodsStatus   string
}

// 初始化AlipayBaoxianClaimReturngoodsstatusUpdateRequest对象
func NewAlipayBaoxianClaimReturngoodsstatusUpdateRequest() *AlipayBaoxianClaimReturngoodsstatusUpdateRequest{
    return &AlipayBaoxianClaimReturngoodsstatusUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlipayBaoxianClaimReturngoodsstatusUpdateRequest) GetApiMethodName() string {
    return "alipay.baoxian.claim.returngoodsstatus.update"
}

// IRequest interface 方法, 获取API参数
func (r AlipayBaoxianClaimReturngoodsstatusUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ClaimNo Setter
// 理赔单号
func (r *AlipayBaoxianClaimReturngoodsstatusUpdateRequest) SetClaimNo(_claimNo string) error {
    r._claimNo = _claimNo
    r.Set("claim_no", _claimNo)
    return nil
}

// ClaimNo Getter
func (r AlipayBaoxianClaimReturngoodsstatusUpdateRequest) GetClaimNo() string {
    return r._claimNo
}
// GoodsStatus Setter
// 退货货物状态
func (r *AlipayBaoxianClaimReturngoodsstatusUpdateRequest) SetGoodsStatus(_goodsStatus string) error {
    r._goodsStatus = _goodsStatus
    r.Set("goods_status", _goodsStatus)
    return nil
}

// GoodsStatus Getter
func (r AlipayBaoxianClaimReturngoodsstatusUpdateRequest) GetGoodsStatus() string {
    return r._goodsStatus
}

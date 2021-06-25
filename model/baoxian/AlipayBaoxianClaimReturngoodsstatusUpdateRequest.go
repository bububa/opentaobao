package baoxian

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
更新理赔单退货货物状态 APIRequest
alipay.baoxian.claim.returngoodsstatus.update

更新理赔单退货货物状态
*/
type AlipayBaoxianClaimReturngoodsstatusUpdateRequest struct {
    model.Params

    // 理赔单号
    claimNo   string 

    // 退货货物状态
    goodsStatus   string 

}

func NewAlipayBaoxianClaimReturngoodsstatusUpdateRequest() *AlipayBaoxianClaimReturngoodsstatusUpdateRequest{
    return &AlipayBaoxianClaimReturngoodsstatusUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r AlipayBaoxianClaimReturngoodsstatusUpdateRequest) GetApiMethodName() string {
    return "alipay.baoxian.claim.returngoodsstatus.update"
}

func (r AlipayBaoxianClaimReturngoodsstatusUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlipayBaoxianClaimReturngoodsstatusUpdateRequest) SetClaimNo(claimNo string) error {
    r.claimNo = claimNo
    r.Set("claim_no", claimNo)
    return nil
}

func (r AlipayBaoxianClaimReturngoodsstatusUpdateRequest) GetClaimNo() string {
    return r.claimNo
}

func (r *AlipayBaoxianClaimReturngoodsstatusUpdateRequest) SetGoodsStatus(goodsStatus string) error {
    r.goodsStatus = goodsStatus
    r.Set("goods_status", goodsStatus)
    return nil
}

func (r AlipayBaoxianClaimReturngoodsstatusUpdateRequest) GetGoodsStatus() string {
    return r.goodsStatus
}


package tmallsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
买家拒收 APIRequest
tmall.servicecenter.workcard.refuse

买家拒收通知接口
*/
type TmallServicecenterWorkcardRefuseRequest struct {
    model.Params

    // 买家拒收信息
    buyerRefuseAcceptRequest   *BuyerRefuseAcceptRequest 

}

func NewTmallServicecenterWorkcardRefuseRequest() *TmallServicecenterWorkcardRefuseRequest{
    return &TmallServicecenterWorkcardRefuseRequest{
        Params: model.NewParams(),
    }
}

func (r TmallServicecenterWorkcardRefuseRequest) GetApiMethodName() string {
    return "tmall.servicecenter.workcard.refuse"
}

func (r TmallServicecenterWorkcardRefuseRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallServicecenterWorkcardRefuseRequest) SetBuyerRefuseAcceptRequest(buyerRefuseAcceptRequest *BuyerRefuseAcceptRequest) error {
    r.buyerRefuseAcceptRequest = buyerRefuseAcceptRequest
    r.Set("buyer_refuse_accept_request", buyerRefuseAcceptRequest)
    return nil
}

func (r TmallServicecenterWorkcardRefuseRequest) GetBuyerRefuseAcceptRequest() *BuyerRefuseAcceptRequest {
    return r.buyerRefuseAcceptRequest
}


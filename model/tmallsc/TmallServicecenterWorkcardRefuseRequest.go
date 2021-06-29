package tmallsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
买家拒收 API请求
tmall.servicecenter.workcard.refuse

买家拒收通知接口
*/
type TmallServicecenterWorkcardRefuseRequest struct {
    model.Params
    // 买家拒收信息
    _buyerRefuseAcceptRequest   *BuyerRefuseAcceptRequest
}

// 初始化TmallServicecenterWorkcardRefuseRequest对象
func NewTmallServicecenterWorkcardRefuseRequest() *TmallServicecenterWorkcardRefuseRequest{
    return &TmallServicecenterWorkcardRefuseRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallServicecenterWorkcardRefuseRequest) GetApiMethodName() string {
    return "tmall.servicecenter.workcard.refuse"
}

// IRequest interface 方法, 获取API参数
func (r TmallServicecenterWorkcardRefuseRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BuyerRefuseAcceptRequest Setter
// 买家拒收信息
func (r *TmallServicecenterWorkcardRefuseRequest) SetBuyerRefuseAcceptRequest(_buyerRefuseAcceptRequest *BuyerRefuseAcceptRequest) error {
    r._buyerRefuseAcceptRequest = _buyerRefuseAcceptRequest
    r.Set("buyer_refuse_accept_request", _buyerRefuseAcceptRequest)
    return nil
}

// BuyerRefuseAcceptRequest Getter
func (r TmallServicecenterWorkcardRefuseRequest) GetBuyerRefuseAcceptRequest() *BuyerRefuseAcceptRequest {
    return r._buyerRefuseAcceptRequest
}

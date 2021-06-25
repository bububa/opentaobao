package servicecenter

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
整车租赁风控模型回调 APIRequest
tmall.car.lease.riskcallback

租赁公司回调风控结果
*/
type TmallCarLeaseRiskcallbackRequest struct {
    model.Params

    // 授信结果
    creditInfo   *CreditInfoTopDto 

}

func NewTmallCarLeaseRiskcallbackRequest() *TmallCarLeaseRiskcallbackRequest{
    return &TmallCarLeaseRiskcallbackRequest{
        Params: model.NewParams(),
    }
}

func (r TmallCarLeaseRiskcallbackRequest) GetApiMethodName() string {
    return "tmall.car.lease.riskcallback"
}

func (r TmallCarLeaseRiskcallbackRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallCarLeaseRiskcallbackRequest) SetCreditInfo(creditInfo *CreditInfoTopDto) error {
    r.creditInfo = creditInfo
    r.Set("credit_info", creditInfo)
    return nil
}

func (r TmallCarLeaseRiskcallbackRequest) GetCreditInfo() *CreditInfoTopDto {
    return r.creditInfo
}


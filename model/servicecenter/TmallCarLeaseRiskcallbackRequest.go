package servicecenter

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
整车租赁风控模型回调 API请求
tmall.car.lease.riskcallback

租赁公司回调风控结果
*/
type TmallCarLeaseRiskcallbackRequest struct {
    model.Params
    // 授信结果
    creditInfo   *CreditInfoTopDto
}

// 初始化TmallCarLeaseRiskcallbackRequest对象
func NewTmallCarLeaseRiskcallbackRequest() *TmallCarLeaseRiskcallbackRequest{
    return &TmallCarLeaseRiskcallbackRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallCarLeaseRiskcallbackRequest) GetApiMethodName() string {
    return "tmall.car.lease.riskcallback"
}

// IRequest interface 方法, 获取API参数
func (r TmallCarLeaseRiskcallbackRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CreditInfo Setter
// 授信结果
func (r *TmallCarLeaseRiskcallbackRequest) SetCreditInfo(creditInfo *CreditInfoTopDto) error {
    r.creditInfo = creditInfo
    r.Set("credit_info", creditInfo)
    return nil
}

// CreditInfo Getter
func (r TmallCarLeaseRiskcallbackRequest) GetCreditInfo() *CreditInfoTopDto {
    return r.creditInfo
}

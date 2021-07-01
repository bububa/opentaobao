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
type TmallCarLeaseRiskcallbackAPIRequest struct {
    model.Params
    // 授信结果
    _creditInfo   *CreditInfoTopDTO
}

// 初始化TmallCarLeaseRiskcallbackAPIRequest对象
func NewTmallCarLeaseRiskcallbackRequest() *TmallCarLeaseRiskcallbackAPIRequest{
    return &TmallCarLeaseRiskcallbackAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallCarLeaseRiskcallbackAPIRequest) GetApiMethodName() string {
    return "tmall.car.lease.riskcallback"
}

// IRequest interface 方法, 获取API参数
func (r TmallCarLeaseRiskcallbackAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CreditInfo Setter
// 授信结果
func (r *TmallCarLeaseRiskcallbackAPIRequest) SetCreditInfo(_creditInfo *CreditInfoTopDTO) error {
    r._creditInfo = _creditInfo
    r.Set("credit_info", _creditInfo)
    return nil
}

// CreditInfo Getter
func (r TmallCarLeaseRiskcallbackAPIRequest) GetCreditInfo() *CreditInfoTopDTO {
    return r._creditInfo
}

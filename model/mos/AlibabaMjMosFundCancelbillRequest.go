package mos

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
取消付款单 API请求
alibaba.mj.mos.fund.cancelbill

取消付款单
*/
type AlibabaMjMosFundCancelbillRequest struct {
    model.Params
    // 取消入参
    _cancelBillDTO   *CancelBillDTO
}

// 初始化AlibabaMjMosFundCancelbillRequest对象
func NewAlibabaMjMosFundCancelbillRequest() *AlibabaMjMosFundCancelbillRequest{
    return &AlibabaMjMosFundCancelbillRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaMjMosFundCancelbillRequest) GetApiMethodName() string {
    return "alibaba.mj.mos.fund.cancelbill"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaMjMosFundCancelbillRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CancelBillDTO Setter
// 取消入参
func (r *AlibabaMjMosFundCancelbillRequest) SetCancelBillDTO(_cancelBillDTO *CancelBillDTO) error {
    r._cancelBillDTO = _cancelBillDTO
    r.Set("cancel_bill_d_t_o", _cancelBillDTO)
    return nil
}

// CancelBillDTO Getter
func (r AlibabaMjMosFundCancelbillRequest) GetCancelBillDTO() *CancelBillDTO {
    return r._cancelBillDTO
}

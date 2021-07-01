package mos

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创建一个付款单 API请求
alibaba.mj.mos.fund.createbill

创建一个付款单
*/
type AlibabaMjMosFundCreatebillAPIRequest struct {
    model.Params
    // 创建付款单入参
    _billDto   *CreateBillDTO
}

// 初始化AlibabaMjMosFundCreatebillAPIRequest对象
func NewAlibabaMjMosFundCreatebillRequest() *AlibabaMjMosFundCreatebillAPIRequest{
    return &AlibabaMjMosFundCreatebillAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaMjMosFundCreatebillAPIRequest) GetApiMethodName() string {
    return "alibaba.mj.mos.fund.createbill"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaMjMosFundCreatebillAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BillDto Setter
// 创建付款单入参
func (r *AlibabaMjMosFundCreatebillAPIRequest) SetBillDto(_billDto *CreateBillDTO) error {
    r._billDto = _billDto
    r.Set("bill_dto", _billDto)
    return nil
}

// BillDto Getter
func (r AlibabaMjMosFundCreatebillAPIRequest) GetBillDto() *CreateBillDTO {
    return r._billDto
}

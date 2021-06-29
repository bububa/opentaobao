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
type AlibabaMjMosFundCreatebillRequest struct {
    model.Params
    // 创建付款单入参
    billDto   *CreateBillDto
}

// 初始化AlibabaMjMosFundCreatebillRequest对象
func NewAlibabaMjMosFundCreatebillRequest() *AlibabaMjMosFundCreatebillRequest{
    return &AlibabaMjMosFundCreatebillRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaMjMosFundCreatebillRequest) GetApiMethodName() string {
    return "alibaba.mj.mos.fund.createbill"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaMjMosFundCreatebillRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BillDto Setter
// 创建付款单入参
func (r *AlibabaMjMosFundCreatebillRequest) SetBillDto(billDto *CreateBillDto) error {
    r.billDto = billDto
    r.Set("bill_dto", billDto)
    return nil
}

// BillDto Getter
func (r AlibabaMjMosFundCreatebillRequest) GetBillDto() *CreateBillDto {
    return r.billDto
}

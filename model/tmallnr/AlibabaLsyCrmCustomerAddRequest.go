package tmallnr

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
私域导购添加活动留资入口 API请求
alibaba.lsy.crm.customer.add

私域导购添加活动留资入口
*/
type AlibabaLsyCrmCustomerAddRequest struct {
    model.Params
    // 入参对象
    reqDto   *NrtCrmCreateCustomerReq
}

// 初始化AlibabaLsyCrmCustomerAddRequest对象
func NewAlibabaLsyCrmCustomerAddRequest() *AlibabaLsyCrmCustomerAddRequest{
    return &AlibabaLsyCrmCustomerAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaLsyCrmCustomerAddRequest) GetApiMethodName() string {
    return "alibaba.lsy.crm.customer.add"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaLsyCrmCustomerAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ReqDto Setter
// 入参对象
func (r *AlibabaLsyCrmCustomerAddRequest) SetReqDto(reqDto *NrtCrmCreateCustomerReq) error {
    r.reqDto = reqDto
    r.Set("req_dto", reqDto)
    return nil
}

// ReqDto Getter
func (r AlibabaLsyCrmCustomerAddRequest) GetReqDto() *NrtCrmCreateCustomerReq {
    return r.reqDto
}

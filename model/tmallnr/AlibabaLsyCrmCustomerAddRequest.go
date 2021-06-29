package tmallnr

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
私域导购添加活动留资入口 APIRequest
alibaba.lsy.crm.customer.add

私域导购添加活动留资入口
*/
type AlibabaLsyCrmCustomerAddRequest struct {
    model.Params

    // 入参对象
    reqDto   *NrtCrmCreateCustomerReq 

}

func NewAlibabaLsyCrmCustomerAddRequest() *AlibabaLsyCrmCustomerAddRequest{
    return &AlibabaLsyCrmCustomerAddRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaLsyCrmCustomerAddRequest) GetApiMethodName() string {
    return "alibaba.lsy.crm.customer.add"
}

func (r AlibabaLsyCrmCustomerAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaLsyCrmCustomerAddRequest) SetReqDto(reqDto *NrtCrmCreateCustomerReq) error {
    r.reqDto = reqDto
    r.Set("req_dto", reqDto)
    return nil
}

func (r AlibabaLsyCrmCustomerAddRequest) GetReqDto() *NrtCrmCreateCustomerReq {
    return r.reqDto
}


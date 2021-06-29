package fundplatform

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据制卡单分页查询卡信息 APIRequest
alibaba.fundplatform.cardorders.info.query

该接口由汇金实现，外部调用。通过制卡单号分页查询卡信息
*/
type AlibabaFundplatformCardordersInfoQueryRequest struct {
    model.Params

    // 请求结构体
    parameters   *CardMakingInfoQueryRequest 

}

func NewAlibabaFundplatformCardordersInfoQueryRequest() *AlibabaFundplatformCardordersInfoQueryRequest{
    return &AlibabaFundplatformCardordersInfoQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaFundplatformCardordersInfoQueryRequest) GetApiMethodName() string {
    return "alibaba.fundplatform.cardorders.info.query"
}

func (r AlibabaFundplatformCardordersInfoQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaFundplatformCardordersInfoQueryRequest) SetParameters(parameters *CardMakingInfoQueryRequest) error {
    r.parameters = parameters
    r.Set("parameters", parameters)
    return nil
}

func (r AlibabaFundplatformCardordersInfoQueryRequest) GetParameters() *CardMakingInfoQueryRequest {
    return r.parameters
}


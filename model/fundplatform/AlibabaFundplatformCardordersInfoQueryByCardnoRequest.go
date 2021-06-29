package fundplatform

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
通过卡号查询卡信息 APIRequest
alibaba.fundplatform.cardorders.info.query.by.cardno

该接口由汇金实现，外部调用。通过制卡单号分页查询卡信息
*/
type AlibabaFundplatformCardordersInfoQueryByCardnoRequest struct {
    model.Params

    // 请求结构体
    parameters   *CardMakingInfoQueryRequest 

}

func NewAlibabaFundplatformCardordersInfoQueryByCardnoRequest() *AlibabaFundplatformCardordersInfoQueryByCardnoRequest{
    return &AlibabaFundplatformCardordersInfoQueryByCardnoRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaFundplatformCardordersInfoQueryByCardnoRequest) GetApiMethodName() string {
    return "alibaba.fundplatform.cardorders.info.query.by.cardno"
}

func (r AlibabaFundplatformCardordersInfoQueryByCardnoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaFundplatformCardordersInfoQueryByCardnoRequest) SetParameters(parameters *CardMakingInfoQueryRequest) error {
    r.parameters = parameters
    r.Set("parameters", parameters)
    return nil
}

func (r AlibabaFundplatformCardordersInfoQueryByCardnoRequest) GetParameters() *CardMakingInfoQueryRequest {
    return r.parameters
}


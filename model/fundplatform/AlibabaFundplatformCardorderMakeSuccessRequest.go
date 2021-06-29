package fundplatform

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
通知制卡成功 APIRequest
alibaba.fundplatform.cardorder.make.success

当外部业务方调用资金平台异步制卡接口后，资金平台制卡成功后通知异步通知业务方
*/
type AlibabaFundplatformCardorderMakeSuccessRequest struct {
    model.Params

    // 入参对象
    request   *AlibabaFundplatformCardorderMakeSuccessStruct 

}

func NewAlibabaFundplatformCardorderMakeSuccessRequest() *AlibabaFundplatformCardorderMakeSuccessRequest{
    return &AlibabaFundplatformCardorderMakeSuccessRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaFundplatformCardorderMakeSuccessRequest) GetApiMethodName() string {
    return "alibaba.fundplatform.cardorder.make.success"
}

func (r AlibabaFundplatformCardorderMakeSuccessRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaFundplatformCardorderMakeSuccessRequest) SetRequest(request *AlibabaFundplatformCardorderMakeSuccessStruct) error {
    r.request = request
    r.Set("request", request)
    return nil
}

func (r AlibabaFundplatformCardorderMakeSuccessRequest) GetRequest() *AlibabaFundplatformCardorderMakeSuccessStruct {
    return r.request
}


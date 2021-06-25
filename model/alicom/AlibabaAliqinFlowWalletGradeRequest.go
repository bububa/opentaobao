package alicom

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取流量档位 APIRequest
alibaba.aliqin.flow.wallet.grade

获取直充流量档位
*/
type AlibabaAliqinFlowWalletGradeRequest struct {
    model.Params

    // 手机号码
    phoneNum   string 

}

func NewAlibabaAliqinFlowWalletGradeRequest() *AlibabaAliqinFlowWalletGradeRequest{
    return &AlibabaAliqinFlowWalletGradeRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAliqinFlowWalletGradeRequest) GetApiMethodName() string {
    return "alibaba.aliqin.flow.wallet.grade"
}

func (r AlibabaAliqinFlowWalletGradeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAliqinFlowWalletGradeRequest) SetPhoneNum(phoneNum string) error {
    r.phoneNum = phoneNum
    r.Set("phone_num", phoneNum)
    return nil
}

func (r AlibabaAliqinFlowWalletGradeRequest) GetPhoneNum() string {
    return r.phoneNum
}


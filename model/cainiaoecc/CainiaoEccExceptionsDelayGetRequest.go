package cainiaoecc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
菜鸟控制塔包裹滞留异常信息获取 APIRequest
cainiao.ecc.exceptions.delay.get

菜鸟控制塔包裹滞留异常信息获取
*/
type CainiaoEccExceptionsDelayGetRequest struct {
    model.Params

    // 运单号
    mailNo   string 

}

func NewCainiaoEccExceptionsDelayGetRequest() *CainiaoEccExceptionsDelayGetRequest{
    return &CainiaoEccExceptionsDelayGetRequest{
        Params: model.NewParams(),
    }
}

func (r CainiaoEccExceptionsDelayGetRequest) GetApiMethodName() string {
    return "cainiao.ecc.exceptions.delay.get"
}

func (r CainiaoEccExceptionsDelayGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *CainiaoEccExceptionsDelayGetRequest) SetMailNo(mailNo string) error {
    r.mailNo = mailNo
    r.Set("mail_no", mailNo)
    return nil
}

func (r CainiaoEccExceptionsDelayGetRequest) GetMailNo() string {
    return r.mailNo
}


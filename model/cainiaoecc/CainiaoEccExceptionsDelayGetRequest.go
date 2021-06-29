package cainiaoecc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
菜鸟控制塔包裹滞留异常信息获取 API请求
cainiao.ecc.exceptions.delay.get

菜鸟控制塔包裹滞留异常信息获取
*/
type CainiaoEccExceptionsDelayGetRequest struct {
    model.Params
    // 运单号
    mailNo   string
}

// 初始化CainiaoEccExceptionsDelayGetRequest对象
func NewCainiaoEccExceptionsDelayGetRequest() *CainiaoEccExceptionsDelayGetRequest{
    return &CainiaoEccExceptionsDelayGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r CainiaoEccExceptionsDelayGetRequest) GetApiMethodName() string {
    return "cainiao.ecc.exceptions.delay.get"
}

// IRequest interface 方法, 获取API参数
func (r CainiaoEccExceptionsDelayGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// MailNo Setter
// 运单号
func (r *CainiaoEccExceptionsDelayGetRequest) SetMailNo(mailNo string) error {
    r.mailNo = mailNo
    r.Set("mail_no", mailNo)
    return nil
}

// MailNo Getter
func (r CainiaoEccExceptionsDelayGetRequest) GetMailNo() string {
    return r.mailNo
}

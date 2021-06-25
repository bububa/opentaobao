package guoguo

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
兜底派送订单的运单号回传接口 APIRequest
cainiao.guoguo.backup.graborder.submitmailno

快递公司回传订单号和运单号给菜鸟裹裹
*/
type CainiaoGuoguoBackupGraborderSubmitmailnoRequest struct {
    model.Params

    // 菜鸟物流订单号
    orderCode   string 

    // 运单号
    mailNo   string 

}

func NewCainiaoGuoguoBackupGraborderSubmitmailnoRequest() *CainiaoGuoguoBackupGraborderSubmitmailnoRequest{
    return &CainiaoGuoguoBackupGraborderSubmitmailnoRequest{
        Params: model.NewParams(),
    }
}

func (r CainiaoGuoguoBackupGraborderSubmitmailnoRequest) GetApiMethodName() string {
    return "cainiao.guoguo.backup.graborder.submitmailno"
}

func (r CainiaoGuoguoBackupGraborderSubmitmailnoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *CainiaoGuoguoBackupGraborderSubmitmailnoRequest) SetOrderCode(orderCode string) error {
    r.orderCode = orderCode
    r.Set("orderCode", orderCode)
    return nil
}

func (r CainiaoGuoguoBackupGraborderSubmitmailnoRequest) GetOrderCode() string {
    return r.orderCode
}

func (r *CainiaoGuoguoBackupGraborderSubmitmailnoRequest) SetMailNo(mailNo string) error {
    r.mailNo = mailNo
    r.Set("mailNo", mailNo)
    return nil
}

func (r CainiaoGuoguoBackupGraborderSubmitmailnoRequest) GetMailNo() string {
    return r.mailNo
}


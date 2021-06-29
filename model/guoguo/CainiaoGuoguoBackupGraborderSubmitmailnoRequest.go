package guoguo

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
兜底派送订单的运单号回传接口 API请求
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

// 初始化CainiaoGuoguoBackupGraborderSubmitmailnoRequest对象
func NewCainiaoGuoguoBackupGraborderSubmitmailnoRequest() *CainiaoGuoguoBackupGraborderSubmitmailnoRequest{
    return &CainiaoGuoguoBackupGraborderSubmitmailnoRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r CainiaoGuoguoBackupGraborderSubmitmailnoRequest) GetApiMethodName() string {
    return "cainiao.guoguo.backup.graborder.submitmailno"
}

// IRequest interface 方法, 获取API参数
func (r CainiaoGuoguoBackupGraborderSubmitmailnoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderCode Setter
// 菜鸟物流订单号
func (r *CainiaoGuoguoBackupGraborderSubmitmailnoRequest) SetOrderCode(orderCode string) error {
    r.orderCode = orderCode
    r.Set("orderCode", orderCode)
    return nil
}

// OrderCode Getter
func (r CainiaoGuoguoBackupGraborderSubmitmailnoRequest) GetOrderCode() string {
    return r.orderCode
}
// MailNo Setter
// 运单号
func (r *CainiaoGuoguoBackupGraborderSubmitmailnoRequest) SetMailNo(mailNo string) error {
    r.mailNo = mailNo
    r.Set("mailNo", mailNo)
    return nil
}

// MailNo Getter
func (r CainiaoGuoguoBackupGraborderSubmitmailnoRequest) GetMailNo() string {
    return r.mailNo
}

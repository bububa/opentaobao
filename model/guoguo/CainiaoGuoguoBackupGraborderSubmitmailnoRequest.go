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
type CainiaoGuoguoBackupGraborderSubmitmailnoAPIRequest struct {
    model.Params
    // 菜鸟物流订单号
    _orderCode   string
    // 运单号
    _mailNo   string
}

// 初始化CainiaoGuoguoBackupGraborderSubmitmailnoAPIRequest对象
func NewCainiaoGuoguoBackupGraborderSubmitmailnoRequest() *CainiaoGuoguoBackupGraborderSubmitmailnoAPIRequest{
    return &CainiaoGuoguoBackupGraborderSubmitmailnoAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r CainiaoGuoguoBackupGraborderSubmitmailnoAPIRequest) GetApiMethodName() string {
    return "cainiao.guoguo.backup.graborder.submitmailno"
}

// IRequest interface 方法, 获取API参数
func (r CainiaoGuoguoBackupGraborderSubmitmailnoAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderCode Setter
// 菜鸟物流订单号
func (r *CainiaoGuoguoBackupGraborderSubmitmailnoAPIRequest) SetOrderCode(_orderCode string) error {
    r._orderCode = _orderCode
    r.Set("orderCode", _orderCode)
    return nil
}

// OrderCode Getter
func (r CainiaoGuoguoBackupGraborderSubmitmailnoAPIRequest) GetOrderCode() string {
    return r._orderCode
}
// MailNo Setter
// 运单号
func (r *CainiaoGuoguoBackupGraborderSubmitmailnoAPIRequest) SetMailNo(_mailNo string) error {
    r._mailNo = _mailNo
    r.Set("mailNo", _mailNo)
    return nil
}

// MailNo Getter
func (r CainiaoGuoguoBackupGraborderSubmitmailnoAPIRequest) GetMailNo() string {
    return r._mailNo
}

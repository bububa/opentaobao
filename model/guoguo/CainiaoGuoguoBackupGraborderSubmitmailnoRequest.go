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
    _orderCode   string
    // 运单号
    _mailNo   string
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
func (r *CainiaoGuoguoBackupGraborderSubmitmailnoRequest) SetOrderCode(_orderCode string) error {
    r._orderCode = _orderCode
    r.Set("orderCode", _orderCode)
    return nil
}

// OrderCode Getter
func (r CainiaoGuoguoBackupGraborderSubmitmailnoRequest) GetOrderCode() string {
    return r._orderCode
}
// MailNo Setter
// 运单号
func (r *CainiaoGuoguoBackupGraborderSubmitmailnoRequest) SetMailNo(_mailNo string) error {
    r._mailNo = _mailNo
    r.Set("mailNo", _mailNo)
    return nil
}

// MailNo Getter
func (r CainiaoGuoguoBackupGraborderSubmitmailnoRequest) GetMailNo() string {
    return r._mailNo
}

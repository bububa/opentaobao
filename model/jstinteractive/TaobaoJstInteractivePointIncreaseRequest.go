package jstinteractive

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
互动积分发放接口 API请求
taobao.jst.interactive.point.increase

向用户发放互动积分
*/
type TaobaoJstInteractivePointIncreaseRequest struct {
    model.Params
    // 发放的积分值
    _amount   int64
    // 幂等操作码，要确保唯一性，同一个操作码只能使用一次，避免重复操作
    _operateCode   string
}

// 初始化TaobaoJstInteractivePointIncreaseRequest对象
func NewTaobaoJstInteractivePointIncreaseRequest() *TaobaoJstInteractivePointIncreaseRequest{
    return &TaobaoJstInteractivePointIncreaseRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoJstInteractivePointIncreaseRequest) GetApiMethodName() string {
    return "taobao.jst.interactive.point.increase"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoJstInteractivePointIncreaseRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Amount Setter
// 发放的积分值
func (r *TaobaoJstInteractivePointIncreaseRequest) SetAmount(_amount int64) error {
    r._amount = _amount
    r.Set("amount", _amount)
    return nil
}

// Amount Getter
func (r TaobaoJstInteractivePointIncreaseRequest) GetAmount() int64 {
    return r._amount
}
// OperateCode Setter
// 幂等操作码，要确保唯一性，同一个操作码只能使用一次，避免重复操作
func (r *TaobaoJstInteractivePointIncreaseRequest) SetOperateCode(_operateCode string) error {
    r._operateCode = _operateCode
    r.Set("operate_code", _operateCode)
    return nil
}

// OperateCode Getter
func (r TaobaoJstInteractivePointIncreaseRequest) GetOperateCode() string {
    return r._operateCode
}

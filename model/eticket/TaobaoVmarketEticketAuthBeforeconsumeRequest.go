package eticket

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
核销放行的查询接口 API请求
taobao.vmarket.eticket.auth.beforeconsume

针对O2O电子凭证核销放行业务，为满足码商能够核销淘宝码而开放的核销查询接口
*/
type TaobaoVmarketEticketAuthBeforeconsumeRequest struct {
    model.Params
    // 核销的码，只支持单个码，多个码核销需要多次调用
    _verifyCode   string
    // 核销方的ID，如果是普通码商必须传入机具ID,如果是私有码商家（即原有的信任商家）可默认传入私有码商ID
    _operatorid   string
    // 网点ID,网点授权核销时，必须传入；其他核销方式可不传
    _storeid   string
}

// 初始化TaobaoVmarketEticketAuthBeforeconsumeRequest对象
func NewTaobaoVmarketEticketAuthBeforeconsumeRequest() *TaobaoVmarketEticketAuthBeforeconsumeRequest{
    return &TaobaoVmarketEticketAuthBeforeconsumeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoVmarketEticketAuthBeforeconsumeRequest) GetApiMethodName() string {
    return "taobao.vmarket.eticket.auth.beforeconsume"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoVmarketEticketAuthBeforeconsumeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// VerifyCode Setter
// 核销的码，只支持单个码，多个码核销需要多次调用
func (r *TaobaoVmarketEticketAuthBeforeconsumeRequest) SetVerifyCode(_verifyCode string) error {
    r._verifyCode = _verifyCode
    r.Set("verify_code", _verifyCode)
    return nil
}

// VerifyCode Getter
func (r TaobaoVmarketEticketAuthBeforeconsumeRequest) GetVerifyCode() string {
    return r._verifyCode
}
// Operatorid Setter
// 核销方的ID，如果是普通码商必须传入机具ID,如果是私有码商家（即原有的信任商家）可默认传入私有码商ID
func (r *TaobaoVmarketEticketAuthBeforeconsumeRequest) SetOperatorid(_operatorid string) error {
    r._operatorid = _operatorid
    r.Set("operatorid", _operatorid)
    return nil
}

// Operatorid Getter
func (r TaobaoVmarketEticketAuthBeforeconsumeRequest) GetOperatorid() string {
    return r._operatorid
}
// Storeid Setter
// 网点ID,网点授权核销时，必须传入；其他核销方式可不传
func (r *TaobaoVmarketEticketAuthBeforeconsumeRequest) SetStoreid(_storeid string) error {
    r._storeid = _storeid
    r.Set("storeid", _storeid)
    return nil
}

// Storeid Getter
func (r TaobaoVmarketEticketAuthBeforeconsumeRequest) GetStoreid() string {
    return r._storeid
}

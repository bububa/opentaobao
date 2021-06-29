package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
购物金卡查询 API请求
taobao.card.expandcard.query

购物金充值信息查询接口，会返回余额等信息。
*/
type TaobaoCardExpandcardQueryRequest struct {
    model.Params
    // 卡使用范围，不传则会查询所有
    _usedScopeCode   string
    // 支付宝accountNo
    _accountNo   string
}

// 初始化TaobaoCardExpandcardQueryRequest对象
func NewTaobaoCardExpandcardQueryRequest() *TaobaoCardExpandcardQueryRequest{
    return &TaobaoCardExpandcardQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoCardExpandcardQueryRequest) GetApiMethodName() string {
    return "taobao.card.expandcard.query"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoCardExpandcardQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// UsedScopeCode Setter
// 卡使用范围，不传则会查询所有
func (r *TaobaoCardExpandcardQueryRequest) SetUsedScopeCode(_usedScopeCode string) error {
    r._usedScopeCode = _usedScopeCode
    r.Set("used_scope_code", _usedScopeCode)
    return nil
}

// UsedScopeCode Getter
func (r TaobaoCardExpandcardQueryRequest) GetUsedScopeCode() string {
    return r._usedScopeCode
}
// AccountNo Setter
// 支付宝accountNo
func (r *TaobaoCardExpandcardQueryRequest) SetAccountNo(_accountNo string) error {
    r._accountNo = _accountNo
    r.Set("account_no", _accountNo)
    return nil
}

// AccountNo Getter
func (r TaobaoCardExpandcardQueryRequest) GetAccountNo() string {
    return r._accountNo
}

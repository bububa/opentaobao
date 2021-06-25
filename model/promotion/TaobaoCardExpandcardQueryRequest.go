package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
购物金卡查询 APIRequest
taobao.card.expandcard.query

购物金充值信息查询接口，会返回余额等信息。
*/
type TaobaoCardExpandcardQueryRequest struct {
    model.Params

    // 卡使用范围，不传则会查询所有
    usedScopeCode   string 

    // 支付宝accountNo
    accountNo   string 

}

func NewTaobaoCardExpandcardQueryRequest() *TaobaoCardExpandcardQueryRequest{
    return &TaobaoCardExpandcardQueryRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoCardExpandcardQueryRequest) GetApiMethodName() string {
    return "taobao.card.expandcard.query"
}

func (r TaobaoCardExpandcardQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoCardExpandcardQueryRequest) SetUsedScopeCode(usedScopeCode string) error {
    r.usedScopeCode = usedScopeCode
    r.Set("used_scope_code", usedScopeCode)
    return nil
}

func (r TaobaoCardExpandcardQueryRequest) GetUsedScopeCode() string {
    return r.usedScopeCode
}

func (r *TaobaoCardExpandcardQueryRequest) SetAccountNo(accountNo string) error {
    r.accountNo = accountNo
    r.Set("account_no", accountNo)
    return nil
}

func (r TaobaoCardExpandcardQueryRequest) GetAccountNo() string {
    return r.accountNo
}


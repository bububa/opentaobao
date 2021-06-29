package uscesl

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
电子价签解绑接口 APIRequest
taobao.uscesl.biz.esl.unbind

电子价签解绑接口
*/
type TaobaoUsceslBizEslUnbindRequest struct {
    model.Params

    // 价签条码
    eslBarCode   string 

    // 价签系统注册的门店storeId
    storeId   int64 

    // 商家标识key
    bizBrandKey   string 

}

func NewTaobaoUsceslBizEslUnbindRequest() *TaobaoUsceslBizEslUnbindRequest{
    return &TaobaoUsceslBizEslUnbindRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoUsceslBizEslUnbindRequest) GetApiMethodName() string {
    return "taobao.uscesl.biz.esl.unbind"
}

func (r TaobaoUsceslBizEslUnbindRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoUsceslBizEslUnbindRequest) SetEslBarCode(eslBarCode string) error {
    r.eslBarCode = eslBarCode
    r.Set("esl_bar_code", eslBarCode)
    return nil
}

func (r TaobaoUsceslBizEslUnbindRequest) GetEslBarCode() string {
    return r.eslBarCode
}

func (r *TaobaoUsceslBizEslUnbindRequest) SetStoreId(storeId int64) error {
    r.storeId = storeId
    r.Set("store_id", storeId)
    return nil
}

func (r TaobaoUsceslBizEslUnbindRequest) GetStoreId() int64 {
    return r.storeId
}

func (r *TaobaoUsceslBizEslUnbindRequest) SetBizBrandKey(bizBrandKey string) error {
    r.bizBrandKey = bizBrandKey
    r.Set("biz_brand_key", bizBrandKey)
    return nil
}

func (r TaobaoUsceslBizEslUnbindRequest) GetBizBrandKey() string {
    return r.bizBrandKey
}


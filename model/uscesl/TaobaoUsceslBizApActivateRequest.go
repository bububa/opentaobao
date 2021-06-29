package uscesl

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
激活AP价签通讯模块 APIRequest
taobao.uscesl.biz.ap.activate

激活AP价签通讯模块
*/
type TaobaoUsceslBizApActivateRequest struct {
    model.Params

    // AP的mac地址
    apMac   string 

    // 门店ID
    storeId   int64 

    // 商家编码
    bizBrandKey   string 

}

func NewTaobaoUsceslBizApActivateRequest() *TaobaoUsceslBizApActivateRequest{
    return &TaobaoUsceslBizApActivateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoUsceslBizApActivateRequest) GetApiMethodName() string {
    return "taobao.uscesl.biz.ap.activate"
}

func (r TaobaoUsceslBizApActivateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoUsceslBizApActivateRequest) SetApMac(apMac string) error {
    r.apMac = apMac
    r.Set("ap_mac", apMac)
    return nil
}

func (r TaobaoUsceslBizApActivateRequest) GetApMac() string {
    return r.apMac
}

func (r *TaobaoUsceslBizApActivateRequest) SetStoreId(storeId int64) error {
    r.storeId = storeId
    r.Set("store_id", storeId)
    return nil
}

func (r TaobaoUsceslBizApActivateRequest) GetStoreId() int64 {
    return r.storeId
}

func (r *TaobaoUsceslBizApActivateRequest) SetBizBrandKey(bizBrandKey string) error {
    r.bizBrandKey = bizBrandKey
    r.Set("biz_brand_key", bizBrandKey)
    return nil
}

func (r TaobaoUsceslBizApActivateRequest) GetBizBrandKey() string {
    return r.bizBrandKey
}


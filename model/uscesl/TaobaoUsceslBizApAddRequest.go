package uscesl

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
新增价签通讯AP设备 APIRequest
taobao.uscesl.biz.ap.add

根据门店和ap的MAC地址新增
*/
type TaobaoUsceslBizApAddRequest struct {
    model.Params

    // AP MAC地址
    apMac   string 

    // 价签系统门店ID
    storeId   int64 

    // 商家code
    bizBrandKey   string 

}

func NewTaobaoUsceslBizApAddRequest() *TaobaoUsceslBizApAddRequest{
    return &TaobaoUsceslBizApAddRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoUsceslBizApAddRequest) GetApiMethodName() string {
    return "taobao.uscesl.biz.ap.add"
}

func (r TaobaoUsceslBizApAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoUsceslBizApAddRequest) SetApMac(apMac string) error {
    r.apMac = apMac
    r.Set("ap_mac", apMac)
    return nil
}

func (r TaobaoUsceslBizApAddRequest) GetApMac() string {
    return r.apMac
}

func (r *TaobaoUsceslBizApAddRequest) SetStoreId(storeId int64) error {
    r.storeId = storeId
    r.Set("store_id", storeId)
    return nil
}

func (r TaobaoUsceslBizApAddRequest) GetStoreId() int64 {
    return r.storeId
}

func (r *TaobaoUsceslBizApAddRequest) SetBizBrandKey(bizBrandKey string) error {
    r.bizBrandKey = bizBrandKey
    r.Set("biz_brand_key", bizBrandKey)
    return nil
}

func (r TaobaoUsceslBizApAddRequest) GetBizBrandKey() string {
    return r.bizBrandKey
}


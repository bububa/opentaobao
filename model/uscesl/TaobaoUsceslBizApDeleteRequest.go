package uscesl

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除价签AP设备 APIRequest
taobao.uscesl.biz.ap.delete

删除价签AP设备
*/
type TaobaoUsceslBizApDeleteRequest struct {
    model.Params

    // ap的MAC地址
    apMac   string 

    // 门店ID
    storeId   int64 

    // 商家CODE
    bizBrandKey   string 

}

func NewTaobaoUsceslBizApDeleteRequest() *TaobaoUsceslBizApDeleteRequest{
    return &TaobaoUsceslBizApDeleteRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoUsceslBizApDeleteRequest) GetApiMethodName() string {
    return "taobao.uscesl.biz.ap.delete"
}

func (r TaobaoUsceslBizApDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoUsceslBizApDeleteRequest) SetApMac(apMac string) error {
    r.apMac = apMac
    r.Set("ap_mac", apMac)
    return nil
}

func (r TaobaoUsceslBizApDeleteRequest) GetApMac() string {
    return r.apMac
}

func (r *TaobaoUsceslBizApDeleteRequest) SetStoreId(storeId int64) error {
    r.storeId = storeId
    r.Set("store_id", storeId)
    return nil
}

func (r TaobaoUsceslBizApDeleteRequest) GetStoreId() int64 {
    return r.storeId
}

func (r *TaobaoUsceslBizApDeleteRequest) SetBizBrandKey(bizBrandKey string) error {
    r.bizBrandKey = bizBrandKey
    r.Set("biz_brand_key", bizBrandKey)
    return nil
}

func (r TaobaoUsceslBizApDeleteRequest) GetBizBrandKey() string {
    return r.bizBrandKey
}


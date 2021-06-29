package uscesl

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除价签AP设备 API请求
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

// 初始化TaobaoUsceslBizApDeleteRequest对象
func NewTaobaoUsceslBizApDeleteRequest() *TaobaoUsceslBizApDeleteRequest{
    return &TaobaoUsceslBizApDeleteRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoUsceslBizApDeleteRequest) GetApiMethodName() string {
    return "taobao.uscesl.biz.ap.delete"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoUsceslBizApDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ApMac Setter
// ap的MAC地址
func (r *TaobaoUsceslBizApDeleteRequest) SetApMac(apMac string) error {
    r.apMac = apMac
    r.Set("ap_mac", apMac)
    return nil
}

// ApMac Getter
func (r TaobaoUsceslBizApDeleteRequest) GetApMac() string {
    return r.apMac
}
// StoreId Setter
// 门店ID
func (r *TaobaoUsceslBizApDeleteRequest) SetStoreId(storeId int64) error {
    r.storeId = storeId
    r.Set("store_id", storeId)
    return nil
}

// StoreId Getter
func (r TaobaoUsceslBizApDeleteRequest) GetStoreId() int64 {
    return r.storeId
}
// BizBrandKey Setter
// 商家CODE
func (r *TaobaoUsceslBizApDeleteRequest) SetBizBrandKey(bizBrandKey string) error {
    r.bizBrandKey = bizBrandKey
    r.Set("biz_brand_key", bizBrandKey)
    return nil
}

// BizBrandKey Getter
func (r TaobaoUsceslBizApDeleteRequest) GetBizBrandKey() string {
    return r.bizBrandKey
}

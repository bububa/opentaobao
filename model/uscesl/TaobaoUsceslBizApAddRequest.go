package uscesl

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
新增价签通讯AP设备 API请求
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

// 初始化TaobaoUsceslBizApAddRequest对象
func NewTaobaoUsceslBizApAddRequest() *TaobaoUsceslBizApAddRequest{
    return &TaobaoUsceslBizApAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoUsceslBizApAddRequest) GetApiMethodName() string {
    return "taobao.uscesl.biz.ap.add"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoUsceslBizApAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ApMac Setter
// AP MAC地址
func (r *TaobaoUsceslBizApAddRequest) SetApMac(apMac string) error {
    r.apMac = apMac
    r.Set("ap_mac", apMac)
    return nil
}

// ApMac Getter
func (r TaobaoUsceslBizApAddRequest) GetApMac() string {
    return r.apMac
}
// StoreId Setter
// 价签系统门店ID
func (r *TaobaoUsceslBizApAddRequest) SetStoreId(storeId int64) error {
    r.storeId = storeId
    r.Set("store_id", storeId)
    return nil
}

// StoreId Getter
func (r TaobaoUsceslBizApAddRequest) GetStoreId() int64 {
    return r.storeId
}
// BizBrandKey Setter
// 商家code
func (r *TaobaoUsceslBizApAddRequest) SetBizBrandKey(bizBrandKey string) error {
    r.bizBrandKey = bizBrandKey
    r.Set("biz_brand_key", bizBrandKey)
    return nil
}

// BizBrandKey Getter
func (r TaobaoUsceslBizApAddRequest) GetBizBrandKey() string {
    return r.bizBrandKey
}

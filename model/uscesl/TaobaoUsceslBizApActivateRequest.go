package uscesl

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
激活AP价签通讯模块 API请求
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

// 初始化TaobaoUsceslBizApActivateRequest对象
func NewTaobaoUsceslBizApActivateRequest() *TaobaoUsceslBizApActivateRequest{
    return &TaobaoUsceslBizApActivateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoUsceslBizApActivateRequest) GetApiMethodName() string {
    return "taobao.uscesl.biz.ap.activate"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoUsceslBizApActivateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ApMac Setter
// AP的mac地址
func (r *TaobaoUsceslBizApActivateRequest) SetApMac(apMac string) error {
    r.apMac = apMac
    r.Set("ap_mac", apMac)
    return nil
}

// ApMac Getter
func (r TaobaoUsceslBizApActivateRequest) GetApMac() string {
    return r.apMac
}
// StoreId Setter
// 门店ID
func (r *TaobaoUsceslBizApActivateRequest) SetStoreId(storeId int64) error {
    r.storeId = storeId
    r.Set("store_id", storeId)
    return nil
}

// StoreId Getter
func (r TaobaoUsceslBizApActivateRequest) GetStoreId() int64 {
    return r.storeId
}
// BizBrandKey Setter
// 商家编码
func (r *TaobaoUsceslBizApActivateRequest) SetBizBrandKey(bizBrandKey string) error {
    r.bizBrandKey = bizBrandKey
    r.Set("biz_brand_key", bizBrandKey)
    return nil
}

// BizBrandKey Getter
func (r TaobaoUsceslBizApActivateRequest) GetBizBrandKey() string {
    return r.bizBrandKey
}

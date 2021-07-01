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
type TaobaoUsceslBizApActivateAPIRequest struct {
    model.Params
    // AP的mac地址
    _apMac   string
    // 门店ID
    _storeId   int64
    // 商家编码
    _bizBrandKey   string
}

// 初始化TaobaoUsceslBizApActivateAPIRequest对象
func NewTaobaoUsceslBizApActivateRequest() *TaobaoUsceslBizApActivateAPIRequest{
    return &TaobaoUsceslBizApActivateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoUsceslBizApActivateAPIRequest) GetApiMethodName() string {
    return "taobao.uscesl.biz.ap.activate"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoUsceslBizApActivateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ApMac Setter
// AP的mac地址
func (r *TaobaoUsceslBizApActivateAPIRequest) SetApMac(_apMac string) error {
    r._apMac = _apMac
    r.Set("ap_mac", _apMac)
    return nil
}

// ApMac Getter
func (r TaobaoUsceslBizApActivateAPIRequest) GetApMac() string {
    return r._apMac
}
// StoreId Setter
// 门店ID
func (r *TaobaoUsceslBizApActivateAPIRequest) SetStoreId(_storeId int64) error {
    r._storeId = _storeId
    r.Set("store_id", _storeId)
    return nil
}

// StoreId Getter
func (r TaobaoUsceslBizApActivateAPIRequest) GetStoreId() int64 {
    return r._storeId
}
// BizBrandKey Setter
// 商家编码
func (r *TaobaoUsceslBizApActivateAPIRequest) SetBizBrandKey(_bizBrandKey string) error {
    r._bizBrandKey = _bizBrandKey
    r.Set("biz_brand_key", _bizBrandKey)
    return nil
}

// BizBrandKey Getter
func (r TaobaoUsceslBizApActivateAPIRequest) GetBizBrandKey() string {
    return r._bizBrandKey
}

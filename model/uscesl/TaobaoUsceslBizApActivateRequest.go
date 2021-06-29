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
    _apMac   string
    // 门店ID
    _storeId   int64
    // 商家编码
    _bizBrandKey   string
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
func (r *TaobaoUsceslBizApActivateRequest) SetApMac(_apMac string) error {
    r._apMac = _apMac
    r.Set("ap_mac", _apMac)
    return nil
}

// ApMac Getter
func (r TaobaoUsceslBizApActivateRequest) GetApMac() string {
    return r._apMac
}
// StoreId Setter
// 门店ID
func (r *TaobaoUsceslBizApActivateRequest) SetStoreId(_storeId int64) error {
    r._storeId = _storeId
    r.Set("store_id", _storeId)
    return nil
}

// StoreId Getter
func (r TaobaoUsceslBizApActivateRequest) GetStoreId() int64 {
    return r._storeId
}
// BizBrandKey Setter
// 商家编码
func (r *TaobaoUsceslBizApActivateRequest) SetBizBrandKey(_bizBrandKey string) error {
    r._bizBrandKey = _bizBrandKey
    r.Set("biz_brand_key", _bizBrandKey)
    return nil
}

// BizBrandKey Getter
func (r TaobaoUsceslBizApActivateRequest) GetBizBrandKey() string {
    return r._bizBrandKey
}

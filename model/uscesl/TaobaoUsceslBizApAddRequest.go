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
    _apMac   string
    // 价签系统门店ID
    _storeId   int64
    // 商家code
    _bizBrandKey   string
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
func (r *TaobaoUsceslBizApAddRequest) SetApMac(_apMac string) error {
    r._apMac = _apMac
    r.Set("ap_mac", _apMac)
    return nil
}

// ApMac Getter
func (r TaobaoUsceslBizApAddRequest) GetApMac() string {
    return r._apMac
}
// StoreId Setter
// 价签系统门店ID
func (r *TaobaoUsceslBizApAddRequest) SetStoreId(_storeId int64) error {
    r._storeId = _storeId
    r.Set("store_id", _storeId)
    return nil
}

// StoreId Getter
func (r TaobaoUsceslBizApAddRequest) GetStoreId() int64 {
    return r._storeId
}
// BizBrandKey Setter
// 商家code
func (r *TaobaoUsceslBizApAddRequest) SetBizBrandKey(_bizBrandKey string) error {
    r._bizBrandKey = _bizBrandKey
    r.Set("biz_brand_key", _bizBrandKey)
    return nil
}

// BizBrandKey Getter
func (r TaobaoUsceslBizApAddRequest) GetBizBrandKey() string {
    return r._bizBrandKey
}

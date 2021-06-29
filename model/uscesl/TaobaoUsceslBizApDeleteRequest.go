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
    _apMac   string
    // 门店ID
    _storeId   int64
    // 商家CODE
    _bizBrandKey   string
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
func (r *TaobaoUsceslBizApDeleteRequest) SetApMac(_apMac string) error {
    r._apMac = _apMac
    r.Set("ap_mac", _apMac)
    return nil
}

// ApMac Getter
func (r TaobaoUsceslBizApDeleteRequest) GetApMac() string {
    return r._apMac
}
// StoreId Setter
// 门店ID
func (r *TaobaoUsceslBizApDeleteRequest) SetStoreId(_storeId int64) error {
    r._storeId = _storeId
    r.Set("store_id", _storeId)
    return nil
}

// StoreId Getter
func (r TaobaoUsceslBizApDeleteRequest) GetStoreId() int64 {
    return r._storeId
}
// BizBrandKey Setter
// 商家CODE
func (r *TaobaoUsceslBizApDeleteRequest) SetBizBrandKey(_bizBrandKey string) error {
    r._bizBrandKey = _bizBrandKey
    r.Set("biz_brand_key", _bizBrandKey)
    return nil
}

// BizBrandKey Getter
func (r TaobaoUsceslBizApDeleteRequest) GetBizBrandKey() string {
    return r._bizBrandKey
}

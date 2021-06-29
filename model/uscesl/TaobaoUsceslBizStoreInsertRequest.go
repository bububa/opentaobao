package uscesl

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
新增电子价签商家门店接口 API请求
taobao.uscesl.biz.store.insert

新增电子价签商家门店接口
*/
type TaobaoUsceslBizStoreInsertRequest struct {
    model.Params
    // 门店名称
    _storeName   string
    // 门店外部ID，要保持同一商家下的唯一性
    _storeOutId   string
    // 商家标识key
    _bizBrandKey   string
}

// 初始化TaobaoUsceslBizStoreInsertRequest对象
func NewTaobaoUsceslBizStoreInsertRequest() *TaobaoUsceslBizStoreInsertRequest{
    return &TaobaoUsceslBizStoreInsertRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoUsceslBizStoreInsertRequest) GetApiMethodName() string {
    return "taobao.uscesl.biz.store.insert"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoUsceslBizStoreInsertRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// StoreName Setter
// 门店名称
func (r *TaobaoUsceslBizStoreInsertRequest) SetStoreName(_storeName string) error {
    r._storeName = _storeName
    r.Set("store_name", _storeName)
    return nil
}

// StoreName Getter
func (r TaobaoUsceslBizStoreInsertRequest) GetStoreName() string {
    return r._storeName
}
// StoreOutId Setter
// 门店外部ID，要保持同一商家下的唯一性
func (r *TaobaoUsceslBizStoreInsertRequest) SetStoreOutId(_storeOutId string) error {
    r._storeOutId = _storeOutId
    r.Set("store_out_id", _storeOutId)
    return nil
}

// StoreOutId Getter
func (r TaobaoUsceslBizStoreInsertRequest) GetStoreOutId() string {
    return r._storeOutId
}
// BizBrandKey Setter
// 商家标识key
func (r *TaobaoUsceslBizStoreInsertRequest) SetBizBrandKey(_bizBrandKey string) error {
    r._bizBrandKey = _bizBrandKey
    r.Set("biz_brand_key", _bizBrandKey)
    return nil
}

// BizBrandKey Getter
func (r TaobaoUsceslBizStoreInsertRequest) GetBizBrandKey() string {
    return r._bizBrandKey
}

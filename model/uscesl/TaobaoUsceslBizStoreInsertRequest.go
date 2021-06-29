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
    storeName   string
    // 门店外部ID，要保持同一商家下的唯一性
    storeOutId   string
    // 商家标识key
    bizBrandKey   string
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
func (r *TaobaoUsceslBizStoreInsertRequest) SetStoreName(storeName string) error {
    r.storeName = storeName
    r.Set("store_name", storeName)
    return nil
}

// StoreName Getter
func (r TaobaoUsceslBizStoreInsertRequest) GetStoreName() string {
    return r.storeName
}
// StoreOutId Setter
// 门店外部ID，要保持同一商家下的唯一性
func (r *TaobaoUsceslBizStoreInsertRequest) SetStoreOutId(storeOutId string) error {
    r.storeOutId = storeOutId
    r.Set("store_out_id", storeOutId)
    return nil
}

// StoreOutId Getter
func (r TaobaoUsceslBizStoreInsertRequest) GetStoreOutId() string {
    return r.storeOutId
}
// BizBrandKey Setter
// 商家标识key
func (r *TaobaoUsceslBizStoreInsertRequest) SetBizBrandKey(bizBrandKey string) error {
    r.bizBrandKey = bizBrandKey
    r.Set("biz_brand_key", bizBrandKey)
    return nil
}

// BizBrandKey Getter
func (r TaobaoUsceslBizStoreInsertRequest) GetBizBrandKey() string {
    return r.bizBrandKey
}

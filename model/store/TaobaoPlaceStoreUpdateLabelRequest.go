package store

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商户门店标签更新接口 API请求
taobao.place.store.update.label

更新商户门店标签（服务、权益、标签）接口
*/
type TaobaoPlaceStoreUpdateLabelRequest struct {
    model.Params
    // 门店id
    _storeId   int64
    // 标签id
    _labelIdList   []int64
    // 行业code
    _businessCode   string
    // 标签类型
    _labelType   string
}

// 初始化TaobaoPlaceStoreUpdateLabelRequest对象
func NewTaobaoPlaceStoreUpdateLabelRequest() *TaobaoPlaceStoreUpdateLabelRequest{
    return &TaobaoPlaceStoreUpdateLabelRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoPlaceStoreUpdateLabelRequest) GetApiMethodName() string {
    return "taobao.place.store.update.label"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoPlaceStoreUpdateLabelRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// StoreId Setter
// 门店id
func (r *TaobaoPlaceStoreUpdateLabelRequest) SetStoreId(_storeId int64) error {
    r._storeId = _storeId
    r.Set("store_id", _storeId)
    return nil
}

// StoreId Getter
func (r TaobaoPlaceStoreUpdateLabelRequest) GetStoreId() int64 {
    return r._storeId
}
// LabelIdList Setter
// 标签id
func (r *TaobaoPlaceStoreUpdateLabelRequest) SetLabelIdList(_labelIdList []int64) error {
    r._labelIdList = _labelIdList
    r.Set("label_id_list", _labelIdList)
    return nil
}

// LabelIdList Getter
func (r TaobaoPlaceStoreUpdateLabelRequest) GetLabelIdList() []int64 {
    return r._labelIdList
}
// BusinessCode Setter
// 行业code
func (r *TaobaoPlaceStoreUpdateLabelRequest) SetBusinessCode(_businessCode string) error {
    r._businessCode = _businessCode
    r.Set("business_code", _businessCode)
    return nil
}

// BusinessCode Getter
func (r TaobaoPlaceStoreUpdateLabelRequest) GetBusinessCode() string {
    return r._businessCode
}
// LabelType Setter
// 标签类型
func (r *TaobaoPlaceStoreUpdateLabelRequest) SetLabelType(_labelType string) error {
    r._labelType = _labelType
    r.Set("label_type", _labelType)
    return nil
}

// LabelType Getter
func (r TaobaoPlaceStoreUpdateLabelRequest) GetLabelType() string {
    return r._labelType
}

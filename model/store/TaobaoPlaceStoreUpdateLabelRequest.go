package store

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商户门店标签更新接口 APIRequest
taobao.place.store.update.label

更新商户门店标签（服务、权益、标签）接口
*/
type TaobaoPlaceStoreUpdateLabelRequest struct {
    model.Params

    // 门店id
    storeId   int64 

    // 标签id
    labelIdList   []int64 

    // 行业code
    businessCode   string 

    // 标签类型
    labelType   string 

}

func NewTaobaoPlaceStoreUpdateLabelRequest() *TaobaoPlaceStoreUpdateLabelRequest{
    return &TaobaoPlaceStoreUpdateLabelRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoPlaceStoreUpdateLabelRequest) GetApiMethodName() string {
    return "taobao.place.store.update.label"
}

func (r TaobaoPlaceStoreUpdateLabelRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoPlaceStoreUpdateLabelRequest) SetStoreId(storeId int64) error {
    r.storeId = storeId
    r.Set("store_id", storeId)
    return nil
}

func (r TaobaoPlaceStoreUpdateLabelRequest) GetStoreId() int64 {
    return r.storeId
}

func (r *TaobaoPlaceStoreUpdateLabelRequest) SetLabelIdList(labelIdList []int64) error {
    r.labelIdList = labelIdList
    r.Set("label_id_list", labelIdList)
    return nil
}

func (r TaobaoPlaceStoreUpdateLabelRequest) GetLabelIdList() []int64 {
    return r.labelIdList
}

func (r *TaobaoPlaceStoreUpdateLabelRequest) SetBusinessCode(businessCode string) error {
    r.businessCode = businessCode
    r.Set("business_code", businessCode)
    return nil
}

func (r TaobaoPlaceStoreUpdateLabelRequest) GetBusinessCode() string {
    return r.businessCode
}

func (r *TaobaoPlaceStoreUpdateLabelRequest) SetLabelType(labelType string) error {
    r.labelType = labelType
    r.Set("label_type", labelType)
    return nil
}

func (r TaobaoPlaceStoreUpdateLabelRequest) GetLabelType() string {
    return r.labelType
}


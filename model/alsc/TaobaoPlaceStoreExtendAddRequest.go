package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
新增门店扩展属性 APIRequest
taobao.place.store.extend.add

新增授权用户的门店扩展属性
*/
type TaobaoPlaceStoreExtendAddRequest struct {
    model.Params

    // 门店ID
    storeId   int64 

    // 扩展信息
    etv   []ExtendTypeValueTopDto 

}

func NewTaobaoPlaceStoreExtendAddRequest() *TaobaoPlaceStoreExtendAddRequest{
    return &TaobaoPlaceStoreExtendAddRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoPlaceStoreExtendAddRequest) GetApiMethodName() string {
    return "taobao.place.store.extend.add"
}

func (r TaobaoPlaceStoreExtendAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoPlaceStoreExtendAddRequest) SetStoreId(storeId int64) error {
    r.storeId = storeId
    r.Set("store_id", storeId)
    return nil
}

func (r TaobaoPlaceStoreExtendAddRequest) GetStoreId() int64 {
    return r.storeId
}

func (r *TaobaoPlaceStoreExtendAddRequest) SetEtv(etv []ExtendTypeValueTopDto) error {
    r.etv = etv
    r.Set("etv", etv)
    return nil
}

func (r TaobaoPlaceStoreExtendAddRequest) GetEtv() []ExtendTypeValueTopDto {
    return r.etv
}


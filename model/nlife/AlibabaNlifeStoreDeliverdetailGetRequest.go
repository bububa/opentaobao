package nlife

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询发货单详情 APIRequest
alibaba.nlife.store.deliverdetail.get

查询发货单详情
*/
type AlibabaNlifeStoreDeliverdetailGetRequest struct {
    model.Params

    // 发货单号
    consignNo   string 

    // 门店id
    storeId   int64 

}

func NewAlibabaNlifeStoreDeliverdetailGetRequest() *AlibabaNlifeStoreDeliverdetailGetRequest{
    return &AlibabaNlifeStoreDeliverdetailGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaNlifeStoreDeliverdetailGetRequest) GetApiMethodName() string {
    return "alibaba.nlife.store.deliverdetail.get"
}

func (r AlibabaNlifeStoreDeliverdetailGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaNlifeStoreDeliverdetailGetRequest) SetConsignNo(consignNo string) error {
    r.consignNo = consignNo
    r.Set("consign_no", consignNo)
    return nil
}

func (r AlibabaNlifeStoreDeliverdetailGetRequest) GetConsignNo() string {
    return r.consignNo
}

func (r *AlibabaNlifeStoreDeliverdetailGetRequest) SetStoreId(storeId int64) error {
    r.storeId = storeId
    r.Set("store_id", storeId)
    return nil
}

func (r AlibabaNlifeStoreDeliverdetailGetRequest) GetStoreId() int64 {
    return r.storeId
}


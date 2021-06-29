package nlife

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询发货单详情 API请求
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

// 初始化AlibabaNlifeStoreDeliverdetailGetRequest对象
func NewAlibabaNlifeStoreDeliverdetailGetRequest() *AlibabaNlifeStoreDeliverdetailGetRequest{
    return &AlibabaNlifeStoreDeliverdetailGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaNlifeStoreDeliverdetailGetRequest) GetApiMethodName() string {
    return "alibaba.nlife.store.deliverdetail.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaNlifeStoreDeliverdetailGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ConsignNo Setter
// 发货单号
func (r *AlibabaNlifeStoreDeliverdetailGetRequest) SetConsignNo(consignNo string) error {
    r.consignNo = consignNo
    r.Set("consign_no", consignNo)
    return nil
}

// ConsignNo Getter
func (r AlibabaNlifeStoreDeliverdetailGetRequest) GetConsignNo() string {
    return r.consignNo
}
// StoreId Setter
// 门店id
func (r *AlibabaNlifeStoreDeliverdetailGetRequest) SetStoreId(storeId int64) error {
    r.storeId = storeId
    r.Set("store_id", storeId)
    return nil
}

// StoreId Getter
func (r AlibabaNlifeStoreDeliverdetailGetRequest) GetStoreId() int64 {
    return r.storeId
}

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
type AlibabaNlifeStoreDeliverdetailGetAPIRequest struct {
    model.Params
    // 发货单号
    _consignNo   string
    // 门店id
    _storeId   int64
}

// 初始化AlibabaNlifeStoreDeliverdetailGetAPIRequest对象
func NewAlibabaNlifeStoreDeliverdetailGetRequest() *AlibabaNlifeStoreDeliverdetailGetAPIRequest{
    return &AlibabaNlifeStoreDeliverdetailGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaNlifeStoreDeliverdetailGetAPIRequest) GetApiMethodName() string {
    return "alibaba.nlife.store.deliverdetail.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaNlifeStoreDeliverdetailGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ConsignNo Setter
// 发货单号
func (r *AlibabaNlifeStoreDeliverdetailGetAPIRequest) SetConsignNo(_consignNo string) error {
    r._consignNo = _consignNo
    r.Set("consign_no", _consignNo)
    return nil
}

// ConsignNo Getter
func (r AlibabaNlifeStoreDeliverdetailGetAPIRequest) GetConsignNo() string {
    return r._consignNo
}
// StoreId Setter
// 门店id
func (r *AlibabaNlifeStoreDeliverdetailGetAPIRequest) SetStoreId(_storeId int64) error {
    r._storeId = _storeId
    r.Set("store_id", _storeId)
    return nil
}

// StoreId Getter
func (r AlibabaNlifeStoreDeliverdetailGetAPIRequest) GetStoreId() int64 {
    return r._storeId
}

package nlife

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取门店采购单下的发货单列表 API请求
alibaba.nlife.store.delivers.get

获取门店采购单下的发货单列表
*/
type AlibabaNlifeStoreDeliversGetRequest struct {
    model.Params
    // 门店采购订单号
    _tradeNo   string
    // 零售商的门店id
    _storeId   int64
    // 每页的数量
    _pageSize   int64
    // 查询的页码
    _pageNo   int64
}

// 初始化AlibabaNlifeStoreDeliversGetRequest对象
func NewAlibabaNlifeStoreDeliversGetRequest() *AlibabaNlifeStoreDeliversGetRequest{
    return &AlibabaNlifeStoreDeliversGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaNlifeStoreDeliversGetRequest) GetApiMethodName() string {
    return "alibaba.nlife.store.delivers.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaNlifeStoreDeliversGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TradeNo Setter
// 门店采购订单号
func (r *AlibabaNlifeStoreDeliversGetRequest) SetTradeNo(_tradeNo string) error {
    r._tradeNo = _tradeNo
    r.Set("trade_no", _tradeNo)
    return nil
}

// TradeNo Getter
func (r AlibabaNlifeStoreDeliversGetRequest) GetTradeNo() string {
    return r._tradeNo
}
// StoreId Setter
// 零售商的门店id
func (r *AlibabaNlifeStoreDeliversGetRequest) SetStoreId(_storeId int64) error {
    r._storeId = _storeId
    r.Set("store_id", _storeId)
    return nil
}

// StoreId Getter
func (r AlibabaNlifeStoreDeliversGetRequest) GetStoreId() int64 {
    return r._storeId
}
// PageSize Setter
// 每页的数量
func (r *AlibabaNlifeStoreDeliversGetRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r AlibabaNlifeStoreDeliversGetRequest) GetPageSize() int64 {
    return r._pageSize
}
// PageNo Setter
// 查询的页码
func (r *AlibabaNlifeStoreDeliversGetRequest) SetPageNo(_pageNo int64) error {
    r._pageNo = _pageNo
    r.Set("page_no", _pageNo)
    return nil
}

// PageNo Getter
func (r AlibabaNlifeStoreDeliversGetRequest) GetPageNo() int64 {
    return r._pageNo
}

package nlife

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
b2c下载订单 API请求
alibaba.nlife.b2c.trade.download

下载零售商在零售+平台创建的订单
*/
type AlibabaNlifeB2cTradeDownloadRequest struct {
    model.Params
    // 页码
    _pageNo   int64
    // 分页大小
    _pageSize   int64
    // 零售门店在零售+平台对应的ID
    _storeId   string
    // 开始时间
    _startDate   string
    // 结束时间
    _endDate   string
}

// 初始化AlibabaNlifeB2cTradeDownloadRequest对象
func NewAlibabaNlifeB2cTradeDownloadRequest() *AlibabaNlifeB2cTradeDownloadRequest{
    return &AlibabaNlifeB2cTradeDownloadRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaNlifeB2cTradeDownloadRequest) GetApiMethodName() string {
    return "alibaba.nlife.b2c.trade.download"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaNlifeB2cTradeDownloadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PageNo Setter
// 页码
func (r *AlibabaNlifeB2cTradeDownloadRequest) SetPageNo(_pageNo int64) error {
    r._pageNo = _pageNo
    r.Set("page_no", _pageNo)
    return nil
}

// PageNo Getter
func (r AlibabaNlifeB2cTradeDownloadRequest) GetPageNo() int64 {
    return r._pageNo
}
// PageSize Setter
// 分页大小
func (r *AlibabaNlifeB2cTradeDownloadRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r AlibabaNlifeB2cTradeDownloadRequest) GetPageSize() int64 {
    return r._pageSize
}
// StoreId Setter
// 零售门店在零售+平台对应的ID
func (r *AlibabaNlifeB2cTradeDownloadRequest) SetStoreId(_storeId string) error {
    r._storeId = _storeId
    r.Set("store_id", _storeId)
    return nil
}

// StoreId Getter
func (r AlibabaNlifeB2cTradeDownloadRequest) GetStoreId() string {
    return r._storeId
}
// StartDate Setter
// 开始时间
func (r *AlibabaNlifeB2cTradeDownloadRequest) SetStartDate(_startDate string) error {
    r._startDate = _startDate
    r.Set("start_date", _startDate)
    return nil
}

// StartDate Getter
func (r AlibabaNlifeB2cTradeDownloadRequest) GetStartDate() string {
    return r._startDate
}
// EndDate Setter
// 结束时间
func (r *AlibabaNlifeB2cTradeDownloadRequest) SetEndDate(_endDate string) error {
    r._endDate = _endDate
    r.Set("end_date", _endDate)
    return nil
}

// EndDate Getter
func (r AlibabaNlifeB2cTradeDownloadRequest) GetEndDate() string {
    return r._endDate
}

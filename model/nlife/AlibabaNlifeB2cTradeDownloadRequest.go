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
    pageNo   int64
    // 分页大小
    pageSize   int64
    // 零售门店在零售+平台对应的ID
    storeId   string
    // 开始时间
    startDate   string
    // 结束时间
    endDate   string
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
func (r *AlibabaNlifeB2cTradeDownloadRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

// PageNo Getter
func (r AlibabaNlifeB2cTradeDownloadRequest) GetPageNo() int64 {
    return r.pageNo
}
// PageSize Setter
// 分页大小
func (r *AlibabaNlifeB2cTradeDownloadRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r AlibabaNlifeB2cTradeDownloadRequest) GetPageSize() int64 {
    return r.pageSize
}
// StoreId Setter
// 零售门店在零售+平台对应的ID
func (r *AlibabaNlifeB2cTradeDownloadRequest) SetStoreId(storeId string) error {
    r.storeId = storeId
    r.Set("store_id", storeId)
    return nil
}

// StoreId Getter
func (r AlibabaNlifeB2cTradeDownloadRequest) GetStoreId() string {
    return r.storeId
}
// StartDate Setter
// 开始时间
func (r *AlibabaNlifeB2cTradeDownloadRequest) SetStartDate(startDate string) error {
    r.startDate = startDate
    r.Set("start_date", startDate)
    return nil
}

// StartDate Getter
func (r AlibabaNlifeB2cTradeDownloadRequest) GetStartDate() string {
    return r.startDate
}
// EndDate Setter
// 结束时间
func (r *AlibabaNlifeB2cTradeDownloadRequest) SetEndDate(endDate string) error {
    r.endDate = endDate
    r.Set("end_date", endDate)
    return nil
}

// EndDate Getter
func (r AlibabaNlifeB2cTradeDownloadRequest) GetEndDate() string {
    return r.endDate
}

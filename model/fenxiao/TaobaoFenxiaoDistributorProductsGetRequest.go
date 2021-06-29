package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
分销商查询产品信息 API请求
taobao.fenxiao.distributor.products.get

分销商查询供应商产品信息
*/
type TaobaoFenxiaoDistributorProductsGetRequest struct {
    model.Params
    // order_by
    _orderBy   string
    // time_type
    _timeType   string
    // 下载状态，默认是未下载；可选值：UNDOWNLOAD：未下载 ；DOWNLOADED：已下载；下载：指将供应商授权的产品发布为店铺新宝贝的过程，下载成功后，分销商需要将新生成的宝贝重新编辑并上架后售卖。
    _downloadStatus   string
    // 分销方式；可选择：AGENT ： 代销；DEALER：经销；DIRECT：直营
    _tradeType   string
    // 结束时间
    _endTime   string
    // 指定查询额外的信息，可选值：skus（sku数据）、images（多图），多个可选值用逗号分割。
    _fields   []string
    // 根据商品ID列表查询，优先级次于产品ID列表，高于其他分页查询条件。如果商品不是分销商品，自动过滤。最大限制20，用逗号分割，例如：“1001,1002,1003,1004,1005”
    _itemIds   []int64
    // 产品线ID
    _productcatId   int64
    // 产品ID列表（最大限制30），用逗号分割，例如：“1001,1002,1003,1004,1005”
    _pids   []int64
    // 开始修改时间
    _startTime   string
    // 页码（大于0的整数，默认1）
    _pageNo   int64
    // 每页记录数（默认20，最大50）
    _pageSize   int64
    // 供应商nick，分页查询时，必传
    _supplierNick   string
}

// 初始化TaobaoFenxiaoDistributorProductsGetRequest对象
func NewTaobaoFenxiaoDistributorProductsGetRequest() *TaobaoFenxiaoDistributorProductsGetRequest{
    return &TaobaoFenxiaoDistributorProductsGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFenxiaoDistributorProductsGetRequest) GetApiMethodName() string {
    return "taobao.fenxiao.distributor.products.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFenxiaoDistributorProductsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderBy Setter
// order_by
func (r *TaobaoFenxiaoDistributorProductsGetRequest) SetOrderBy(_orderBy string) error {
    r._orderBy = _orderBy
    r.Set("order_by", _orderBy)
    return nil
}

// OrderBy Getter
func (r TaobaoFenxiaoDistributorProductsGetRequest) GetOrderBy() string {
    return r._orderBy
}
// TimeType Setter
// time_type
func (r *TaobaoFenxiaoDistributorProductsGetRequest) SetTimeType(_timeType string) error {
    r._timeType = _timeType
    r.Set("time_type", _timeType)
    return nil
}

// TimeType Getter
func (r TaobaoFenxiaoDistributorProductsGetRequest) GetTimeType() string {
    return r._timeType
}
// DownloadStatus Setter
// 下载状态，默认是未下载；可选值：UNDOWNLOAD：未下载 ；DOWNLOADED：已下载；下载：指将供应商授权的产品发布为店铺新宝贝的过程，下载成功后，分销商需要将新生成的宝贝重新编辑并上架后售卖。
func (r *TaobaoFenxiaoDistributorProductsGetRequest) SetDownloadStatus(_downloadStatus string) error {
    r._downloadStatus = _downloadStatus
    r.Set("download_status", _downloadStatus)
    return nil
}

// DownloadStatus Getter
func (r TaobaoFenxiaoDistributorProductsGetRequest) GetDownloadStatus() string {
    return r._downloadStatus
}
// TradeType Setter
// 分销方式；可选择：AGENT ： 代销；DEALER：经销；DIRECT：直营
func (r *TaobaoFenxiaoDistributorProductsGetRequest) SetTradeType(_tradeType string) error {
    r._tradeType = _tradeType
    r.Set("trade_type", _tradeType)
    return nil
}

// TradeType Getter
func (r TaobaoFenxiaoDistributorProductsGetRequest) GetTradeType() string {
    return r._tradeType
}
// EndTime Setter
// 结束时间
func (r *TaobaoFenxiaoDistributorProductsGetRequest) SetEndTime(_endTime string) error {
    r._endTime = _endTime
    r.Set("end_time", _endTime)
    return nil
}

// EndTime Getter
func (r TaobaoFenxiaoDistributorProductsGetRequest) GetEndTime() string {
    return r._endTime
}
// Fields Setter
// 指定查询额外的信息，可选值：skus（sku数据）、images（多图），多个可选值用逗号分割。
func (r *TaobaoFenxiaoDistributorProductsGetRequest) SetFields(_fields []string) error {
    r._fields = _fields
    r.Set("fields", _fields)
    return nil
}

// Fields Getter
func (r TaobaoFenxiaoDistributorProductsGetRequest) GetFields() []string {
    return r._fields
}
// ItemIds Setter
// 根据商品ID列表查询，优先级次于产品ID列表，高于其他分页查询条件。如果商品不是分销商品，自动过滤。最大限制20，用逗号分割，例如：“1001,1002,1003,1004,1005”
func (r *TaobaoFenxiaoDistributorProductsGetRequest) SetItemIds(_itemIds []int64) error {
    r._itemIds = _itemIds
    r.Set("item_ids", _itemIds)
    return nil
}

// ItemIds Getter
func (r TaobaoFenxiaoDistributorProductsGetRequest) GetItemIds() []int64 {
    return r._itemIds
}
// ProductcatId Setter
// 产品线ID
func (r *TaobaoFenxiaoDistributorProductsGetRequest) SetProductcatId(_productcatId int64) error {
    r._productcatId = _productcatId
    r.Set("productcat_id", _productcatId)
    return nil
}

// ProductcatId Getter
func (r TaobaoFenxiaoDistributorProductsGetRequest) GetProductcatId() int64 {
    return r._productcatId
}
// Pids Setter
// 产品ID列表（最大限制30），用逗号分割，例如：“1001,1002,1003,1004,1005”
func (r *TaobaoFenxiaoDistributorProductsGetRequest) SetPids(_pids []int64) error {
    r._pids = _pids
    r.Set("pids", _pids)
    return nil
}

// Pids Getter
func (r TaobaoFenxiaoDistributorProductsGetRequest) GetPids() []int64 {
    return r._pids
}
// StartTime Setter
// 开始修改时间
func (r *TaobaoFenxiaoDistributorProductsGetRequest) SetStartTime(_startTime string) error {
    r._startTime = _startTime
    r.Set("start_time", _startTime)
    return nil
}

// StartTime Getter
func (r TaobaoFenxiaoDistributorProductsGetRequest) GetStartTime() string {
    return r._startTime
}
// PageNo Setter
// 页码（大于0的整数，默认1）
func (r *TaobaoFenxiaoDistributorProductsGetRequest) SetPageNo(_pageNo int64) error {
    r._pageNo = _pageNo
    r.Set("page_no", _pageNo)
    return nil
}

// PageNo Getter
func (r TaobaoFenxiaoDistributorProductsGetRequest) GetPageNo() int64 {
    return r._pageNo
}
// PageSize Setter
// 每页记录数（默认20，最大50）
func (r *TaobaoFenxiaoDistributorProductsGetRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoFenxiaoDistributorProductsGetRequest) GetPageSize() int64 {
    return r._pageSize
}
// SupplierNick Setter
// 供应商nick，分页查询时，必传
func (r *TaobaoFenxiaoDistributorProductsGetRequest) SetSupplierNick(_supplierNick string) error {
    r._supplierNick = _supplierNick
    r.Set("supplier_nick", _supplierNick)
    return nil
}

// SupplierNick Getter
func (r TaobaoFenxiaoDistributorProductsGetRequest) GetSupplierNick() string {
    return r._supplierNick
}

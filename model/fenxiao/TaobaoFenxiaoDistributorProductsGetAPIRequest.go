package fenxiao

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFenxiaoDistributorProductsGetAPIRequest 分销商查询产品信息 API请求
// taobao.fenxiao.distributor.products.get
//
// 分销商查询供应商产品信息
type TaobaoFenxiaoDistributorProductsGetAPIRequest struct {
	model.Params
	// 产品ID列表（最大限制30），用逗号分割，例如：“1001,1002,1003,1004,1005”
	_pids []int64
	// 根据商品ID列表查询，优先级次于产品ID列表，高于其他分页查询条件。如果商品不是分销商品，自动过滤。最大限制20，用逗号分割，例如：“1001,1002,1003,1004,1005”
	_itemIds []int64
	// 指定查询额外的信息，可选值：skus（sku数据）、images（多图），多个可选值用逗号分割。
	_fields []string
	// 供应商nick，分页查询时，必传
	_supplierNick string
	// 分销方式；可选择：AGENT ： 代销；DEALER：经销；DIRECT：直营
	_tradeType string
	// 下载状态，默认是未下载；可选值：UNDOWNLOAD：未下载 ；DOWNLOADED：已下载；下载：指将供应商授权的产品发布为店铺新宝贝的过程，下载成功后，分销商需要将新生成的宝贝重新编辑并上架后售卖。
	_downloadStatus string
	// 开始修改时间
	_startTime string
	// 结束时间
	_endTime string
	// time_type
	_timeType string
	// order_by
	_orderBy string
	// 产品线ID
	_productcatId int64
	// 页码（大于0的整数，默认1）
	_pageNo int64
	// 每页记录数（默认20，最大50）
	_pageSize int64
}

// NewTaobaoFenxiaoDistributorProductsGetRequest 初始化TaobaoFenxiaoDistributorProductsGetAPIRequest对象
func NewTaobaoFenxiaoDistributorProductsGetRequest() *TaobaoFenxiaoDistributorProductsGetAPIRequest {
	return &TaobaoFenxiaoDistributorProductsGetAPIRequest{
		Params: model.NewParams(13),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoFenxiaoDistributorProductsGetAPIRequest) Reset() {
	r._pids = r._pids[:0]
	r._itemIds = r._itemIds[:0]
	r._fields = r._fields[:0]
	r._supplierNick = ""
	r._tradeType = ""
	r._downloadStatus = ""
	r._startTime = ""
	r._endTime = ""
	r._timeType = ""
	r._orderBy = ""
	r._productcatId = 0
	r._pageNo = 0
	r._pageSize = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFenxiaoDistributorProductsGetAPIRequest) GetApiMethodName() string {
	return "taobao.fenxiao.distributor.products.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFenxiaoDistributorProductsGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoFenxiaoDistributorProductsGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPids is Pids Setter
// 产品ID列表（最大限制30），用逗号分割，例如：“1001,1002,1003,1004,1005”
func (r *TaobaoFenxiaoDistributorProductsGetAPIRequest) SetPids(_pids []int64) error {
	r._pids = _pids
	r.Set("pids", _pids)
	return nil
}

// GetPids Pids Getter
func (r TaobaoFenxiaoDistributorProductsGetAPIRequest) GetPids() []int64 {
	return r._pids
}

// SetItemIds is ItemIds Setter
// 根据商品ID列表查询，优先级次于产品ID列表，高于其他分页查询条件。如果商品不是分销商品，自动过滤。最大限制20，用逗号分割，例如：“1001,1002,1003,1004,1005”
func (r *TaobaoFenxiaoDistributorProductsGetAPIRequest) SetItemIds(_itemIds []int64) error {
	r._itemIds = _itemIds
	r.Set("item_ids", _itemIds)
	return nil
}

// GetItemIds ItemIds Getter
func (r TaobaoFenxiaoDistributorProductsGetAPIRequest) GetItemIds() []int64 {
	return r._itemIds
}

// SetFields is Fields Setter
// 指定查询额外的信息，可选值：skus（sku数据）、images（多图），多个可选值用逗号分割。
func (r *TaobaoFenxiaoDistributorProductsGetAPIRequest) SetFields(_fields []string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// GetFields Fields Getter
func (r TaobaoFenxiaoDistributorProductsGetAPIRequest) GetFields() []string {
	return r._fields
}

// SetSupplierNick is SupplierNick Setter
// 供应商nick，分页查询时，必传
func (r *TaobaoFenxiaoDistributorProductsGetAPIRequest) SetSupplierNick(_supplierNick string) error {
	r._supplierNick = _supplierNick
	r.Set("supplier_nick", _supplierNick)
	return nil
}

// GetSupplierNick SupplierNick Getter
func (r TaobaoFenxiaoDistributorProductsGetAPIRequest) GetSupplierNick() string {
	return r._supplierNick
}

// SetTradeType is TradeType Setter
// 分销方式；可选择：AGENT ： 代销；DEALER：经销；DIRECT：直营
func (r *TaobaoFenxiaoDistributorProductsGetAPIRequest) SetTradeType(_tradeType string) error {
	r._tradeType = _tradeType
	r.Set("trade_type", _tradeType)
	return nil
}

// GetTradeType TradeType Getter
func (r TaobaoFenxiaoDistributorProductsGetAPIRequest) GetTradeType() string {
	return r._tradeType
}

// SetDownloadStatus is DownloadStatus Setter
// 下载状态，默认是未下载；可选值：UNDOWNLOAD：未下载 ；DOWNLOADED：已下载；下载：指将供应商授权的产品发布为店铺新宝贝的过程，下载成功后，分销商需要将新生成的宝贝重新编辑并上架后售卖。
func (r *TaobaoFenxiaoDistributorProductsGetAPIRequest) SetDownloadStatus(_downloadStatus string) error {
	r._downloadStatus = _downloadStatus
	r.Set("download_status", _downloadStatus)
	return nil
}

// GetDownloadStatus DownloadStatus Getter
func (r TaobaoFenxiaoDistributorProductsGetAPIRequest) GetDownloadStatus() string {
	return r._downloadStatus
}

// SetStartTime is StartTime Setter
// 开始修改时间
func (r *TaobaoFenxiaoDistributorProductsGetAPIRequest) SetStartTime(_startTime string) error {
	r._startTime = _startTime
	r.Set("start_time", _startTime)
	return nil
}

// GetStartTime StartTime Getter
func (r TaobaoFenxiaoDistributorProductsGetAPIRequest) GetStartTime() string {
	return r._startTime
}

// SetEndTime is EndTime Setter
// 结束时间
func (r *TaobaoFenxiaoDistributorProductsGetAPIRequest) SetEndTime(_endTime string) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// GetEndTime EndTime Getter
func (r TaobaoFenxiaoDistributorProductsGetAPIRequest) GetEndTime() string {
	return r._endTime
}

// SetTimeType is TimeType Setter
// time_type
func (r *TaobaoFenxiaoDistributorProductsGetAPIRequest) SetTimeType(_timeType string) error {
	r._timeType = _timeType
	r.Set("time_type", _timeType)
	return nil
}

// GetTimeType TimeType Getter
func (r TaobaoFenxiaoDistributorProductsGetAPIRequest) GetTimeType() string {
	return r._timeType
}

// SetOrderBy is OrderBy Setter
// order_by
func (r *TaobaoFenxiaoDistributorProductsGetAPIRequest) SetOrderBy(_orderBy string) error {
	r._orderBy = _orderBy
	r.Set("order_by", _orderBy)
	return nil
}

// GetOrderBy OrderBy Getter
func (r TaobaoFenxiaoDistributorProductsGetAPIRequest) GetOrderBy() string {
	return r._orderBy
}

// SetProductcatId is ProductcatId Setter
// 产品线ID
func (r *TaobaoFenxiaoDistributorProductsGetAPIRequest) SetProductcatId(_productcatId int64) error {
	r._productcatId = _productcatId
	r.Set("productcat_id", _productcatId)
	return nil
}

// GetProductcatId ProductcatId Getter
func (r TaobaoFenxiaoDistributorProductsGetAPIRequest) GetProductcatId() int64 {
	return r._productcatId
}

// SetPageNo is PageNo Setter
// 页码（大于0的整数，默认1）
func (r *TaobaoFenxiaoDistributorProductsGetAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaoFenxiaoDistributorProductsGetAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetPageSize is PageSize Setter
// 每页记录数（默认20，最大50）
func (r *TaobaoFenxiaoDistributorProductsGetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoFenxiaoDistributorProductsGetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

var poolTaobaoFenxiaoDistributorProductsGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoFenxiaoDistributorProductsGetRequest()
	},
}

// GetTaobaoFenxiaoDistributorProductsGetRequest 从 sync.Pool 获取 TaobaoFenxiaoDistributorProductsGetAPIRequest
func GetTaobaoFenxiaoDistributorProductsGetAPIRequest() *TaobaoFenxiaoDistributorProductsGetAPIRequest {
	return poolTaobaoFenxiaoDistributorProductsGetAPIRequest.Get().(*TaobaoFenxiaoDistributorProductsGetAPIRequest)
}

// ReleaseTaobaoFenxiaoDistributorProductsGetAPIRequest 将 TaobaoFenxiaoDistributorProductsGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoFenxiaoDistributorProductsGetAPIRequest(v *TaobaoFenxiaoDistributorProductsGetAPIRequest) {
	v.Reset()
	poolTaobaoFenxiaoDistributorProductsGetAPIRequest.Put(v)
}

package lstfundbill

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLstTradeOrderFundbillQueryAPIRequest 结算明细数据查询（品牌商视角） API请求
// alibaba.lst.trade.order.fundbill.query
//
// 按照指定日期提供交易账单维度的结算明细数据，比供应商工作台上的结算账单还多一些数据项。
type AlibabaLstTradeOrderFundbillQueryAPIRequest struct {
	model.Params
	// 账单日期，格式：yyyy-MM-dd
	_billDate string
	// 每页最大记录数
	_size int64
	// 页码
	_page int64
	// 为true时,返回相应的商品详细信息，item_id和unit
	_needItemDetail bool
}

// NewAlibabaLstTradeOrderFundbillQueryRequest 初始化AlibabaLstTradeOrderFundbillQueryAPIRequest对象
func NewAlibabaLstTradeOrderFundbillQueryRequest() *AlibabaLstTradeOrderFundbillQueryAPIRequest {
	return &AlibabaLstTradeOrderFundbillQueryAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaLstTradeOrderFundbillQueryAPIRequest) Reset() {
	r._billDate = ""
	r._size = 0
	r._page = 0
	r._needItemDetail = false
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLstTradeOrderFundbillQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.lst.trade.order.fundbill.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLstTradeOrderFundbillQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaLstTradeOrderFundbillQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBillDate is BillDate Setter
// 账单日期，格式：yyyy-MM-dd
func (r *AlibabaLstTradeOrderFundbillQueryAPIRequest) SetBillDate(_billDate string) error {
	r._billDate = _billDate
	r.Set("bill_date", _billDate)
	return nil
}

// GetBillDate BillDate Getter
func (r AlibabaLstTradeOrderFundbillQueryAPIRequest) GetBillDate() string {
	return r._billDate
}

// SetSize is Size Setter
// 每页最大记录数
func (r *AlibabaLstTradeOrderFundbillQueryAPIRequest) SetSize(_size int64) error {
	r._size = _size
	r.Set("size", _size)
	return nil
}

// GetSize Size Getter
func (r AlibabaLstTradeOrderFundbillQueryAPIRequest) GetSize() int64 {
	return r._size
}

// SetPage is Page Setter
// 页码
func (r *AlibabaLstTradeOrderFundbillQueryAPIRequest) SetPage(_page int64) error {
	r._page = _page
	r.Set("page", _page)
	return nil
}

// GetPage Page Getter
func (r AlibabaLstTradeOrderFundbillQueryAPIRequest) GetPage() int64 {
	return r._page
}

// SetNeedItemDetail is NeedItemDetail Setter
// 为true时,返回相应的商品详细信息，item_id和unit
func (r *AlibabaLstTradeOrderFundbillQueryAPIRequest) SetNeedItemDetail(_needItemDetail bool) error {
	r._needItemDetail = _needItemDetail
	r.Set("need_item_detail", _needItemDetail)
	return nil
}

// GetNeedItemDetail NeedItemDetail Getter
func (r AlibabaLstTradeOrderFundbillQueryAPIRequest) GetNeedItemDetail() bool {
	return r._needItemDetail
}

var poolAlibabaLstTradeOrderFundbillQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaLstTradeOrderFundbillQueryRequest()
	},
}

// GetAlibabaLstTradeOrderFundbillQueryRequest 从 sync.Pool 获取 AlibabaLstTradeOrderFundbillQueryAPIRequest
func GetAlibabaLstTradeOrderFundbillQueryAPIRequest() *AlibabaLstTradeOrderFundbillQueryAPIRequest {
	return poolAlibabaLstTradeOrderFundbillQueryAPIRequest.Get().(*AlibabaLstTradeOrderFundbillQueryAPIRequest)
}

// ReleaseAlibabaLstTradeOrderFundbillQueryAPIRequest 将 AlibabaLstTradeOrderFundbillQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaLstTradeOrderFundbillQueryAPIRequest(v *AlibabaLstTradeOrderFundbillQueryAPIRequest) {
	v.Reset()
	poolAlibabaLstTradeOrderFundbillQueryAPIRequest.Put(v)
}

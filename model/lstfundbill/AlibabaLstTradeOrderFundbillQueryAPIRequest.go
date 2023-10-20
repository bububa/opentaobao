package lstfundbill

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabalsttradeorderfundbillqueryAPIRequest 结算明细数据查询（品牌商视角） API请求
// alibaba.lst.trade.order.fundbill.query
//
// 按照指定日期提供交易账单维度的结算明细数据，比供应商工作台上的结算账单还多一些数据项。
type AlibabalsttradeorderfundbillqueryAPIRequest struct {
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

// NewAlibabalsttradeorderfundbillqueryRequest 初始化AlibabalsttradeorderfundbillqueryAPIRequest对象
func NewAlibabalsttradeorderfundbillqueryRequest() *AlibabalsttradeorderfundbillqueryAPIRequest {
	return &AlibabalsttradeorderfundbillqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabalsttradeorderfundbillqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.lst.trade.order.fundbill.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabalsttradeorderfundbillqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabalsttradeorderfundbillqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBillDate is BillDate Setter
// 账单日期，格式：yyyy-MM-dd
func (r *AlibabalsttradeorderfundbillqueryAPIRequest) SetBillDate(_billDate string) error {
	r._billDate = _billDate
	r.Set("bill_date", _billDate)
	return nil
}

// GetBillDate BillDate Getter
func (r AlibabalsttradeorderfundbillqueryAPIRequest) GetBillDate() string {
	return r._billDate
}

// SetSize is Size Setter
// 每页最大记录数
func (r *AlibabalsttradeorderfundbillqueryAPIRequest) SetSize(_size int64) error {
	r._size = _size
	r.Set("size", _size)
	return nil
}

// GetSize Size Getter
func (r AlibabalsttradeorderfundbillqueryAPIRequest) GetSize() int64 {
	return r._size
}

// SetPage is Page Setter
// 页码
func (r *AlibabalsttradeorderfundbillqueryAPIRequest) SetPage(_page int64) error {
	r._page = _page
	r.Set("page", _page)
	return nil
}

// GetPage Page Getter
func (r AlibabalsttradeorderfundbillqueryAPIRequest) GetPage() int64 {
	return r._page
}

// SetNeedItemDetail is NeedItemDetail Setter
// 为true时,返回相应的商品详细信息，item_id和unit
func (r *AlibabalsttradeorderfundbillqueryAPIRequest) SetNeedItemDetail(_needItemDetail bool) error {
	r._needItemDetail = _needItemDetail
	r.Set("need_item_detail", _needItemDetail)
	return nil
}

// GetNeedItemDetail NeedItemDetail Getter
func (r AlibabalsttradeorderfundbillqueryAPIRequest) GetNeedItemDetail() bool {
	return r._needItemDetail
}

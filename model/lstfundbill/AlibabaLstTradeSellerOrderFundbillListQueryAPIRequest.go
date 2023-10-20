package lstfundbill

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabalsttradesellerorderfundbilllistqueryAPIRequest 结算明细数据查询（卖家视角） API请求
// alibaba.lst.trade.seller.order.fundbill.list.query
//
// 提供For供应商的结算接口，以交易账单维度提供开放数据，满足供应商自动化结算的诉求
type AlibabalsttradesellerorderfundbilllistqueryAPIRequest struct {
	model.Params
	// 账单日期，格式：yyyy-MM-dd
	_billDate string
	// 每页最大主订单数，注意：返回的content_list数据按照子订单维度展开
	_size int64
	// 页码
	_page int64
}

// NewAlibabalsttradesellerorderfundbilllistqueryRequest 初始化AlibabalsttradesellerorderfundbilllistqueryAPIRequest对象
func NewAlibabalsttradesellerorderfundbilllistqueryRequest() *AlibabalsttradesellerorderfundbilllistqueryAPIRequest {
	return &AlibabalsttradesellerorderfundbilllistqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabalsttradesellerorderfundbilllistqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.lst.trade.seller.order.fundbill.list.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabalsttradesellerorderfundbilllistqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabalsttradesellerorderfundbilllistqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBillDate is BillDate Setter
// 账单日期，格式：yyyy-MM-dd
func (r *AlibabalsttradesellerorderfundbilllistqueryAPIRequest) SetBillDate(_billDate string) error {
	r._billDate = _billDate
	r.Set("bill_date", _billDate)
	return nil
}

// GetBillDate BillDate Getter
func (r AlibabalsttradesellerorderfundbilllistqueryAPIRequest) GetBillDate() string {
	return r._billDate
}

// SetSize is Size Setter
// 每页最大主订单数，注意：返回的content_list数据按照子订单维度展开
func (r *AlibabalsttradesellerorderfundbilllistqueryAPIRequest) SetSize(_size int64) error {
	r._size = _size
	r.Set("size", _size)
	return nil
}

// GetSize Size Getter
func (r AlibabalsttradesellerorderfundbilllistqueryAPIRequest) GetSize() int64 {
	return r._size
}

// SetPage is Page Setter
// 页码
func (r *AlibabalsttradesellerorderfundbilllistqueryAPIRequest) SetPage(_page int64) error {
	r._page = _page
	r.Set("page", _page)
	return nil
}

// GetPage Page Getter
func (r AlibabalsttradesellerorderfundbilllistqueryAPIRequest) GetPage() int64 {
	return r._page
}

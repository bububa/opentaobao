package lstfundbill

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLstTradeSellerOrderFundbillListQueryAPIRequest 结算明细数据查询（卖家视角） API请求
// alibaba.lst.trade.seller.order.fundbill.list.query
//
// 提供For供应商的结算接口，以交易账单维度提供开放数据，满足供应商自动化结算的诉求
type AlibabaLstTradeSellerOrderFundbillListQueryAPIRequest struct {
	model.Params
	// 每页最大主订单数，注意：返回的content_list数据按照子订单维度展开
	_size int64
	// 账单日期，格式：yyyy-MM-dd
	_billDate string
	// 页码
	_page int64
}

// NewAlibabaLstTradeSellerOrderFundbillListQueryRequest 初始化AlibabaLstTradeSellerOrderFundbillListQueryAPIRequest对象
func NewAlibabaLstTradeSellerOrderFundbillListQueryRequest() *AlibabaLstTradeSellerOrderFundbillListQueryAPIRequest {
	return &AlibabaLstTradeSellerOrderFundbillListQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLstTradeSellerOrderFundbillListQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.lst.trade.seller.order.fundbill.list.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLstTradeSellerOrderFundbillListQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetSize is Size Setter
// 每页最大主订单数，注意：返回的content_list数据按照子订单维度展开
func (r *AlibabaLstTradeSellerOrderFundbillListQueryAPIRequest) SetSize(_size int64) error {
	r._size = _size
	r.Set("size", _size)
	return nil
}

// GetSize Size Getter
func (r AlibabaLstTradeSellerOrderFundbillListQueryAPIRequest) GetSize() int64 {
	return r._size
}

// SetBillDate is BillDate Setter
// 账单日期，格式：yyyy-MM-dd
func (r *AlibabaLstTradeSellerOrderFundbillListQueryAPIRequest) SetBillDate(_billDate string) error {
	r._billDate = _billDate
	r.Set("bill_date", _billDate)
	return nil
}

// GetBillDate BillDate Getter
func (r AlibabaLstTradeSellerOrderFundbillListQueryAPIRequest) GetBillDate() string {
	return r._billDate
}

// SetPage is Page Setter
// 页码
func (r *AlibabaLstTradeSellerOrderFundbillListQueryAPIRequest) SetPage(_page int64) error {
	r._page = _page
	r.Set("page", _page)
	return nil
}

// GetPage Page Getter
func (r AlibabaLstTradeSellerOrderFundbillListQueryAPIRequest) GetPage() int64 {
	return r._page
}

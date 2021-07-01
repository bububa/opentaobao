package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthTracecodesellerBillResultSearchAPIRequest
查询出入库单处理结果api API请求
alibaba.alihealth.tracecodeseller.bill.result.search

查询出入库单处理结果api */
type AlibabaAlihealthTracecodesellerBillResultSearchAPIRequest struct {
	model.Params
	// top身份认证
	_skeyCode string
	// 商家id
	_entInfoId int64
	// 单据编号
	_billCode string
	// 查询开始日期
	_beginDate string
	// 查询结束日期
	_endDate string
	// 不需要
	_sellerName string
	// 每页条数
	_pageSize int64
	// 当前页
	_page int64
}

// NewAlibabaAlihealthTracecodesellerBillResultSearchRequest 初始化AlibabaAlihealthTracecodesellerBillResultSearchAPIRequest对象
func NewAlibabaAlihealthTracecodesellerBillResultSearchRequest() *AlibabaAlihealthTracecodesellerBillResultSearchAPIRequest {
	return &AlibabaAlihealthTracecodesellerBillResultSearchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthTracecodesellerBillResultSearchAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.tracecodeseller.bill.result.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthTracecodesellerBillResultSearchAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is SkeyCode Setter
// top身份认证
func (r *AlibabaAlihealthTracecodesellerBillResultSearchAPIRequest) SetSkeyCode(_skeyCode string) error {
	r._skeyCode = _skeyCode
	r.Set("skey_code", _skeyCode)
	return nil
}

// Get SkeyCode Getter
func (r AlibabaAlihealthTracecodesellerBillResultSearchAPIRequest) GetSkeyCode() string {
	return r._skeyCode
}

// Set is EntInfoId Setter
// 商家id
func (r *AlibabaAlihealthTracecodesellerBillResultSearchAPIRequest) SetEntInfoId(_entInfoId int64) error {
	r._entInfoId = _entInfoId
	r.Set("ent_info_id", _entInfoId)
	return nil
}

// Get EntInfoId Getter
func (r AlibabaAlihealthTracecodesellerBillResultSearchAPIRequest) GetEntInfoId() int64 {
	return r._entInfoId
}

// Set is BillCode Setter
// 单据编号
func (r *AlibabaAlihealthTracecodesellerBillResultSearchAPIRequest) SetBillCode(_billCode string) error {
	r._billCode = _billCode
	r.Set("bill_code", _billCode)
	return nil
}

// Get BillCode Getter
func (r AlibabaAlihealthTracecodesellerBillResultSearchAPIRequest) GetBillCode() string {
	return r._billCode
}

// Set is BeginDate Setter
// 查询开始日期
func (r *AlibabaAlihealthTracecodesellerBillResultSearchAPIRequest) SetBeginDate(_beginDate string) error {
	r._beginDate = _beginDate
	r.Set("begin_date", _beginDate)
	return nil
}

// Get BeginDate Getter
func (r AlibabaAlihealthTracecodesellerBillResultSearchAPIRequest) GetBeginDate() string {
	return r._beginDate
}

// Set is EndDate Setter
// 查询结束日期
func (r *AlibabaAlihealthTracecodesellerBillResultSearchAPIRequest) SetEndDate(_endDate string) error {
	r._endDate = _endDate
	r.Set("end_date", _endDate)
	return nil
}

// Get EndDate Getter
func (r AlibabaAlihealthTracecodesellerBillResultSearchAPIRequest) GetEndDate() string {
	return r._endDate
}

// Set is SellerName Setter
// 不需要
func (r *AlibabaAlihealthTracecodesellerBillResultSearchAPIRequest) SetSellerName(_sellerName string) error {
	r._sellerName = _sellerName
	r.Set("seller_name", _sellerName)
	return nil
}

// Get SellerName Getter
func (r AlibabaAlihealthTracecodesellerBillResultSearchAPIRequest) GetSellerName() string {
	return r._sellerName
}

// Set is PageSize Setter
// 每页条数
func (r *AlibabaAlihealthTracecodesellerBillResultSearchAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// Get PageSize Getter
func (r AlibabaAlihealthTracecodesellerBillResultSearchAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// Set is Page Setter
// 当前页
func (r *AlibabaAlihealthTracecodesellerBillResultSearchAPIRequest) SetPage(_page int64) error {
	r._page = _page
	r.Set("page", _page)
	return nil
}

// Get Page Getter
func (r AlibabaAlihealthTracecodesellerBillResultSearchAPIRequest) GetPage() int64 {
	return r._page
}

package lstfundbill

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
结算明细数据查询（品牌商视角） API请求
alibaba.lst.trade.order.fundbill.query

按照指定日期提供交易账单维度的结算明细数据，比供应商工作台上的结算账单还多一些数据项。
*/
type AlibabaLstTradeOrderFundbillQueryRequest struct {
    model.Params
    // 每页最大记录数
    _size   int64
    // 账单日期，格式：yyyy-MM-dd
    _billDate   string
    // 页码
    _page   int64
    // 为true时,返回相应的商品详细信息，item_id和unit
    _needItemDetail   bool
}

// 初始化AlibabaLstTradeOrderFundbillQueryRequest对象
func NewAlibabaLstTradeOrderFundbillQueryRequest() *AlibabaLstTradeOrderFundbillQueryRequest{
    return &AlibabaLstTradeOrderFundbillQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaLstTradeOrderFundbillQueryRequest) GetApiMethodName() string {
    return "alibaba.lst.trade.order.fundbill.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaLstTradeOrderFundbillQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Size Setter
// 每页最大记录数
func (r *AlibabaLstTradeOrderFundbillQueryRequest) SetSize(_size int64) error {
    r._size = _size
    r.Set("size", _size)
    return nil
}

// Size Getter
func (r AlibabaLstTradeOrderFundbillQueryRequest) GetSize() int64 {
    return r._size
}
// BillDate Setter
// 账单日期，格式：yyyy-MM-dd
func (r *AlibabaLstTradeOrderFundbillQueryRequest) SetBillDate(_billDate string) error {
    r._billDate = _billDate
    r.Set("bill_date", _billDate)
    return nil
}

// BillDate Getter
func (r AlibabaLstTradeOrderFundbillQueryRequest) GetBillDate() string {
    return r._billDate
}
// Page Setter
// 页码
func (r *AlibabaLstTradeOrderFundbillQueryRequest) SetPage(_page int64) error {
    r._page = _page
    r.Set("page", _page)
    return nil
}

// Page Getter
func (r AlibabaLstTradeOrderFundbillQueryRequest) GetPage() int64 {
    return r._page
}
// NeedItemDetail Setter
// 为true时,返回相应的商品详细信息，item_id和unit
func (r *AlibabaLstTradeOrderFundbillQueryRequest) SetNeedItemDetail(_needItemDetail bool) error {
    r._needItemDetail = _needItemDetail
    r.Set("need_item_detail", _needItemDetail)
    return nil
}

// NeedItemDetail Getter
func (r AlibabaLstTradeOrderFundbillQueryRequest) GetNeedItemDetail() bool {
    return r._needItemDetail
}

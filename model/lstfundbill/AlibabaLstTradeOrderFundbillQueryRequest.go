package lstfundbill

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
结算明细数据查询（品牌商视角） APIRequest
alibaba.lst.trade.order.fundbill.query

按照指定日期提供交易账单维度的结算明细数据，比供应商工作台上的结算账单还多一些数据项。
*/
type AlibabaLstTradeOrderFundbillQueryRequest struct {
    model.Params

    // 每页最大记录数
    size   int64 

    // 账单日期，格式：yyyy-MM-dd
    billDate   string 

    // 页码
    page   int64 

    // 为true时,返回相应的商品详细信息，item_id和unit
    needItemDetail   bool 

}

func NewAlibabaLstTradeOrderFundbillQueryRequest() *AlibabaLstTradeOrderFundbillQueryRequest{
    return &AlibabaLstTradeOrderFundbillQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaLstTradeOrderFundbillQueryRequest) GetApiMethodName() string {
    return "alibaba.lst.trade.order.fundbill.query"
}

func (r AlibabaLstTradeOrderFundbillQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaLstTradeOrderFundbillQueryRequest) SetSize(size int64) error {
    r.size = size
    r.Set("size", size)
    return nil
}

func (r AlibabaLstTradeOrderFundbillQueryRequest) GetSize() int64 {
    return r.size
}

func (r *AlibabaLstTradeOrderFundbillQueryRequest) SetBillDate(billDate string) error {
    r.billDate = billDate
    r.Set("bill_date", billDate)
    return nil
}

func (r AlibabaLstTradeOrderFundbillQueryRequest) GetBillDate() string {
    return r.billDate
}

func (r *AlibabaLstTradeOrderFundbillQueryRequest) SetPage(page int64) error {
    r.page = page
    r.Set("page", page)
    return nil
}

func (r AlibabaLstTradeOrderFundbillQueryRequest) GetPage() int64 {
    return r.page
}

func (r *AlibabaLstTradeOrderFundbillQueryRequest) SetNeedItemDetail(needItemDetail bool) error {
    r.needItemDetail = needItemDetail
    r.Set("need_item_detail", needItemDetail)
    return nil
}

func (r AlibabaLstTradeOrderFundbillQueryRequest) GetNeedItemDetail() bool {
    return r.needItemDetail
}


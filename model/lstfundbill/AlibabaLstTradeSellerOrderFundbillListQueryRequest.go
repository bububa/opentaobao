package lstfundbill

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
结算明细数据查询（卖家视角） APIRequest
alibaba.lst.trade.seller.order.fundbill.list.query

提供For供应商的结算接口，以交易账单维度提供开放数据，满足供应商自动化结算的诉求
*/
type AlibabaLstTradeSellerOrderFundbillListQueryRequest struct {
    model.Params

    // 每页最大主订单数，注意：返回的content_list数据按照子订单维度展开
    size   int64 

    // 账单日期，格式：yyyy-MM-dd
    billDate   string 

    // 页码
    page   int64 

}

func NewAlibabaLstTradeSellerOrderFundbillListQueryRequest() *AlibabaLstTradeSellerOrderFundbillListQueryRequest{
    return &AlibabaLstTradeSellerOrderFundbillListQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaLstTradeSellerOrderFundbillListQueryRequest) GetApiMethodName() string {
    return "alibaba.lst.trade.seller.order.fundbill.list.query"
}

func (r AlibabaLstTradeSellerOrderFundbillListQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaLstTradeSellerOrderFundbillListQueryRequest) SetSize(size int64) error {
    r.size = size
    r.Set("size", size)
    return nil
}

func (r AlibabaLstTradeSellerOrderFundbillListQueryRequest) GetSize() int64 {
    return r.size
}

func (r *AlibabaLstTradeSellerOrderFundbillListQueryRequest) SetBillDate(billDate string) error {
    r.billDate = billDate
    r.Set("bill_date", billDate)
    return nil
}

func (r AlibabaLstTradeSellerOrderFundbillListQueryRequest) GetBillDate() string {
    return r.billDate
}

func (r *AlibabaLstTradeSellerOrderFundbillListQueryRequest) SetPage(page int64) error {
    r.page = page
    r.Set("page", page)
    return nil
}

func (r AlibabaLstTradeSellerOrderFundbillListQueryRequest) GetPage() int64 {
    return r.page
}


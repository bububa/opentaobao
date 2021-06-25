package wlbimports

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
物流订单获取 APIRequest
taobao.wlb.imports.order.get

一般进口物流订单获取
*/
type TaobaoWlbImportsOrderGetRequest struct {
    model.Params

    // 交易订单号
    tradeId   int64 

    // 交易订单开始创建时间
    gmtCreateStart   string 

    // 交易订单结束创建时间
    gmtCreateEnd   string 

    // 物流订单状态编码。以下依（物流订单状态编码，描述）的形式列举出来：(TIN_CONSING,发货中),(SENT_WAIT_COMPANY_MAKE_SURE,待仓库收货),(ORDER_CANCELED,订单关闭),(COMPANY_MAKE_SURE,海外仓已揽收),(REJECTED_RECEIVED_BY_COMPANY,海外仓拒收),(ORDER_REFUNDING,退货中),(ORDER_REFUND_BY_COMPANY,订单已退货),(EXCEPTION_IN_COMPANY,海外仓内异常),(FAILED_PAID_SHIPPING_FEE,支付失败),(PAID_SHIPPING_FEE,待仓库发货),(COMPANY_CONSIGN_CONFIRM,海外仓已发货),(WAIT_CUSTOMS_MAKE_SURE,清关已收货),(EXCEPTION_IN_CUSTOMS,清关异常),(CUSTOMSING,清关中),(COMPANY_DELIVERY,国内配送)。
    statusCode   string 

    // 页码。取值范围:大于零的整数; 默认值:1
    pageNo   int64 

    // 每页条数。取值范围:大于0小于等于100的整数; 默认值:40; 最小值：1；最大值:20
    pageSize   int64 

}

func NewTaobaoWlbImportsOrderGetRequest() *TaobaoWlbImportsOrderGetRequest{
    return &TaobaoWlbImportsOrderGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoWlbImportsOrderGetRequest) GetApiMethodName() string {
    return "taobao.wlb.imports.order.get"
}

func (r TaobaoWlbImportsOrderGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoWlbImportsOrderGetRequest) SetTradeId(tradeId int64) error {
    r.tradeId = tradeId
    r.Set("trade_id", tradeId)
    return nil
}

func (r TaobaoWlbImportsOrderGetRequest) GetTradeId() int64 {
    return r.tradeId
}

func (r *TaobaoWlbImportsOrderGetRequest) SetGmtCreateStart(gmtCreateStart string) error {
    r.gmtCreateStart = gmtCreateStart
    r.Set("gmt_create_start", gmtCreateStart)
    return nil
}

func (r TaobaoWlbImportsOrderGetRequest) GetGmtCreateStart() string {
    return r.gmtCreateStart
}

func (r *TaobaoWlbImportsOrderGetRequest) SetGmtCreateEnd(gmtCreateEnd string) error {
    r.gmtCreateEnd = gmtCreateEnd
    r.Set("gmt_create_end", gmtCreateEnd)
    return nil
}

func (r TaobaoWlbImportsOrderGetRequest) GetGmtCreateEnd() string {
    return r.gmtCreateEnd
}

func (r *TaobaoWlbImportsOrderGetRequest) SetStatusCode(statusCode string) error {
    r.statusCode = statusCode
    r.Set("status_code", statusCode)
    return nil
}

func (r TaobaoWlbImportsOrderGetRequest) GetStatusCode() string {
    return r.statusCode
}

func (r *TaobaoWlbImportsOrderGetRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

func (r TaobaoWlbImportsOrderGetRequest) GetPageNo() int64 {
    return r.pageNo
}

func (r *TaobaoWlbImportsOrderGetRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r TaobaoWlbImportsOrderGetRequest) GetPageSize() int64 {
    return r.pageSize
}


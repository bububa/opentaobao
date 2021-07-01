package wlbimports

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
物流订单获取 API请求
taobao.wlb.imports.order.get

一般进口物流订单获取
*/
type TaobaoWlbImportsOrderGetAPIRequest struct {
    model.Params
    // 交易订单号
    _tradeId   int64
    // 交易订单开始创建时间
    _gmtCreateStart   string
    // 交易订单结束创建时间
    _gmtCreateEnd   string
    // 物流订单状态编码。以下依（物流订单状态编码，描述）的形式列举出来：(TIN_CONSING,发货中),(SENT_WAIT_COMPANY_MAKE_SURE,待仓库收货),(ORDER_CANCELED,订单关闭),(COMPANY_MAKE_SURE,海外仓已揽收),(REJECTED_RECEIVED_BY_COMPANY,海外仓拒收),(ORDER_REFUNDING,退货中),(ORDER_REFUND_BY_COMPANY,订单已退货),(EXCEPTION_IN_COMPANY,海外仓内异常),(FAILED_PAID_SHIPPING_FEE,支付失败),(PAID_SHIPPING_FEE,待仓库发货),(COMPANY_CONSIGN_CONFIRM,海外仓已发货),(WAIT_CUSTOMS_MAKE_SURE,清关已收货),(EXCEPTION_IN_CUSTOMS,清关异常),(CUSTOMSING,清关中),(COMPANY_DELIVERY,国内配送)。
    _statusCode   string
    // 页码。取值范围:大于零的整数; 默认值:1
    _pageNo   int64
    // 每页条数。取值范围:大于0小于等于100的整数; 默认值:40; 最小值：1；最大值:20
    _pageSize   int64
}

// 初始化TaobaoWlbImportsOrderGetAPIRequest对象
func NewTaobaoWlbImportsOrderGetRequest() *TaobaoWlbImportsOrderGetAPIRequest{
    return &TaobaoWlbImportsOrderGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoWlbImportsOrderGetAPIRequest) GetApiMethodName() string {
    return "taobao.wlb.imports.order.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoWlbImportsOrderGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TradeId Setter
// 交易订单号
func (r *TaobaoWlbImportsOrderGetAPIRequest) SetTradeId(_tradeId int64) error {
    r._tradeId = _tradeId
    r.Set("trade_id", _tradeId)
    return nil
}

// TradeId Getter
func (r TaobaoWlbImportsOrderGetAPIRequest) GetTradeId() int64 {
    return r._tradeId
}
// GmtCreateStart Setter
// 交易订单开始创建时间
func (r *TaobaoWlbImportsOrderGetAPIRequest) SetGmtCreateStart(_gmtCreateStart string) error {
    r._gmtCreateStart = _gmtCreateStart
    r.Set("gmt_create_start", _gmtCreateStart)
    return nil
}

// GmtCreateStart Getter
func (r TaobaoWlbImportsOrderGetAPIRequest) GetGmtCreateStart() string {
    return r._gmtCreateStart
}
// GmtCreateEnd Setter
// 交易订单结束创建时间
func (r *TaobaoWlbImportsOrderGetAPIRequest) SetGmtCreateEnd(_gmtCreateEnd string) error {
    r._gmtCreateEnd = _gmtCreateEnd
    r.Set("gmt_create_end", _gmtCreateEnd)
    return nil
}

// GmtCreateEnd Getter
func (r TaobaoWlbImportsOrderGetAPIRequest) GetGmtCreateEnd() string {
    return r._gmtCreateEnd
}
// StatusCode Setter
// 物流订单状态编码。以下依（物流订单状态编码，描述）的形式列举出来：(TIN_CONSING,发货中),(SENT_WAIT_COMPANY_MAKE_SURE,待仓库收货),(ORDER_CANCELED,订单关闭),(COMPANY_MAKE_SURE,海外仓已揽收),(REJECTED_RECEIVED_BY_COMPANY,海外仓拒收),(ORDER_REFUNDING,退货中),(ORDER_REFUND_BY_COMPANY,订单已退货),(EXCEPTION_IN_COMPANY,海外仓内异常),(FAILED_PAID_SHIPPING_FEE,支付失败),(PAID_SHIPPING_FEE,待仓库发货),(COMPANY_CONSIGN_CONFIRM,海外仓已发货),(WAIT_CUSTOMS_MAKE_SURE,清关已收货),(EXCEPTION_IN_CUSTOMS,清关异常),(CUSTOMSING,清关中),(COMPANY_DELIVERY,国内配送)。
func (r *TaobaoWlbImportsOrderGetAPIRequest) SetStatusCode(_statusCode string) error {
    r._statusCode = _statusCode
    r.Set("status_code", _statusCode)
    return nil
}

// StatusCode Getter
func (r TaobaoWlbImportsOrderGetAPIRequest) GetStatusCode() string {
    return r._statusCode
}
// PageNo Setter
// 页码。取值范围:大于零的整数; 默认值:1
func (r *TaobaoWlbImportsOrderGetAPIRequest) SetPageNo(_pageNo int64) error {
    r._pageNo = _pageNo
    r.Set("page_no", _pageNo)
    return nil
}

// PageNo Getter
func (r TaobaoWlbImportsOrderGetAPIRequest) GetPageNo() int64 {
    return r._pageNo
}
// PageSize Setter
// 每页条数。取值范围:大于0小于等于100的整数; 默认值:40; 最小值：1；最大值:20
func (r *TaobaoWlbImportsOrderGetAPIRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoWlbImportsOrderGetAPIRequest) GetPageSize() int64 {
    return r._pageSize
}

package jst

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
异常订单信息获取 API请求
taobao.oc.tradetrace.alerts.get

提供订单预警模块的数据查询接口
*/
type TaobaoOcTradetraceAlertsGetRequest struct {
    model.Params
    // 异常类型代码：发货回写淘宝异常=1，商家系统漏单提醒=2，发货超时提醒=3，物流寄送超时=4，买家签收超时=5，物流中转异常=6，退货寄回异常=7，订单追回提醒=8"。
    _abnormalType   int64
    // 返回数据的页码
    _pageNo   int64
    // 返回数据每页包含的记录数目
    _pageSize   int64
}

// 初始化TaobaoOcTradetraceAlertsGetRequest对象
func NewTaobaoOcTradetraceAlertsGetRequest() *TaobaoOcTradetraceAlertsGetRequest{
    return &TaobaoOcTradetraceAlertsGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOcTradetraceAlertsGetRequest) GetApiMethodName() string {
    return "taobao.oc.tradetrace.alerts.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOcTradetraceAlertsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AbnormalType Setter
// 异常类型代码：发货回写淘宝异常=1，商家系统漏单提醒=2，发货超时提醒=3，物流寄送超时=4，买家签收超时=5，物流中转异常=6，退货寄回异常=7，订单追回提醒=8"。
func (r *TaobaoOcTradetraceAlertsGetRequest) SetAbnormalType(_abnormalType int64) error {
    r._abnormalType = _abnormalType
    r.Set("abnormal_type", _abnormalType)
    return nil
}

// AbnormalType Getter
func (r TaobaoOcTradetraceAlertsGetRequest) GetAbnormalType() int64 {
    return r._abnormalType
}
// PageNo Setter
// 返回数据的页码
func (r *TaobaoOcTradetraceAlertsGetRequest) SetPageNo(_pageNo int64) error {
    r._pageNo = _pageNo
    r.Set("page_no", _pageNo)
    return nil
}

// PageNo Getter
func (r TaobaoOcTradetraceAlertsGetRequest) GetPageNo() int64 {
    return r._pageNo
}
// PageSize Setter
// 返回数据每页包含的记录数目
func (r *TaobaoOcTradetraceAlertsGetRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoOcTradetraceAlertsGetRequest) GetPageSize() int64 {
    return r._pageSize
}

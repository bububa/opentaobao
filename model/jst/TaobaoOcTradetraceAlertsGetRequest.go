package jst

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
异常订单信息获取 APIRequest
taobao.oc.tradetrace.alerts.get

提供订单预警模块的数据查询接口
*/
type TaobaoOcTradetraceAlertsGetRequest struct {
    model.Params

    // 异常类型代码：发货回写淘宝异常=1，商家系统漏单提醒=2，发货超时提醒=3，物流寄送超时=4，买家签收超时=5，物流中转异常=6，退货寄回异常=7，订单追回提醒=8"。
    abnormalType   int64 

    // 返回数据的页码
    pageNo   int64 

    // 返回数据每页包含的记录数目
    pageSize   int64 

}

func NewTaobaoOcTradetraceAlertsGetRequest() *TaobaoOcTradetraceAlertsGetRequest{
    return &TaobaoOcTradetraceAlertsGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoOcTradetraceAlertsGetRequest) GetApiMethodName() string {
    return "taobao.oc.tradetrace.alerts.get"
}

func (r TaobaoOcTradetraceAlertsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoOcTradetraceAlertsGetRequest) SetAbnormalType(abnormalType int64) error {
    r.abnormalType = abnormalType
    r.Set("abnormal_type", abnormalType)
    return nil
}

func (r TaobaoOcTradetraceAlertsGetRequest) GetAbnormalType() int64 {
    return r.abnormalType
}

func (r *TaobaoOcTradetraceAlertsGetRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

func (r TaobaoOcTradetraceAlertsGetRequest) GetPageNo() int64 {
    return r.pageNo
}

func (r *TaobaoOcTradetraceAlertsGetRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r TaobaoOcTradetraceAlertsGetRequest) GetPageSize() int64 {
    return r.pageSize
}


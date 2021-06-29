package omniorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
Pos端门店发货 API请求
taobao.omniorder.store.consigned

ISV Pos端门店发货，通知星盘
*/
type TaobaoOmniorderStoreConsignedRequest struct {
    model.Params
    // 跟踪Id
    traceId   string
    // 卖家联系人地址库ID，可以通过taobao.logistics.address.search接口查询到地址库ID。如果为空，取的卖家的默认取货地址
    senderId   int64
    // 家装物流-安装收货人信息,如果为空,则取默认收货人信息
    insReceiverTo   *JzReceiverDto
    // 子订单列表
    subOrderList   []StoreConsignedResult
    // 家装物流-发货参数
    jzTopArgs   *JzTopArgsDto
    // 家装物流-安装公司信息,需要安装时,才填写
    insTpDto   *TpDto
    // 家装物流-家装收货人信息,如果为空,则取默认收货信息
    jzReceiverTo   *JzReceiverDto
    // 淘宝交易主订单ID
    tid   int64
    // ISV系统上报时间
    reportTimestamp   int64
    // 家装物流-物流公司信息
    lgTpDto   *TpDto
}

// 初始化TaobaoOmniorderStoreConsignedRequest对象
func NewTaobaoOmniorderStoreConsignedRequest() *TaobaoOmniorderStoreConsignedRequest{
    return &TaobaoOmniorderStoreConsignedRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOmniorderStoreConsignedRequest) GetApiMethodName() string {
    return "taobao.omniorder.store.consigned"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOmniorderStoreConsignedRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TraceId Setter
// 跟踪Id
func (r *TaobaoOmniorderStoreConsignedRequest) SetTraceId(traceId string) error {
    r.traceId = traceId
    r.Set("trace_id", traceId)
    return nil
}

// TraceId Getter
func (r TaobaoOmniorderStoreConsignedRequest) GetTraceId() string {
    return r.traceId
}
// SenderId Setter
// 卖家联系人地址库ID，可以通过taobao.logistics.address.search接口查询到地址库ID。如果为空，取的卖家的默认取货地址
func (r *TaobaoOmniorderStoreConsignedRequest) SetSenderId(senderId int64) error {
    r.senderId = senderId
    r.Set("sender_id", senderId)
    return nil
}

// SenderId Getter
func (r TaobaoOmniorderStoreConsignedRequest) GetSenderId() int64 {
    return r.senderId
}
// InsReceiverTo Setter
// 家装物流-安装收货人信息,如果为空,则取默认收货人信息
func (r *TaobaoOmniorderStoreConsignedRequest) SetInsReceiverTo(insReceiverTo *JzReceiverDto) error {
    r.insReceiverTo = insReceiverTo
    r.Set("ins_receiver_to", insReceiverTo)
    return nil
}

// InsReceiverTo Getter
func (r TaobaoOmniorderStoreConsignedRequest) GetInsReceiverTo() *JzReceiverDto {
    return r.insReceiverTo
}
// SubOrderList Setter
// 子订单列表
func (r *TaobaoOmniorderStoreConsignedRequest) SetSubOrderList(subOrderList []StoreConsignedResult) error {
    r.subOrderList = subOrderList
    r.Set("sub_order_list", subOrderList)
    return nil
}

// SubOrderList Getter
func (r TaobaoOmniorderStoreConsignedRequest) GetSubOrderList() []StoreConsignedResult {
    return r.subOrderList
}
// JzTopArgs Setter
// 家装物流-发货参数
func (r *TaobaoOmniorderStoreConsignedRequest) SetJzTopArgs(jzTopArgs *JzTopArgsDto) error {
    r.jzTopArgs = jzTopArgs
    r.Set("jz_top_args", jzTopArgs)
    return nil
}

// JzTopArgs Getter
func (r TaobaoOmniorderStoreConsignedRequest) GetJzTopArgs() *JzTopArgsDto {
    return r.jzTopArgs
}
// InsTpDto Setter
// 家装物流-安装公司信息,需要安装时,才填写
func (r *TaobaoOmniorderStoreConsignedRequest) SetInsTpDto(insTpDto *TpDto) error {
    r.insTpDto = insTpDto
    r.Set("ins_tp_dto", insTpDto)
    return nil
}

// InsTpDto Getter
func (r TaobaoOmniorderStoreConsignedRequest) GetInsTpDto() *TpDto {
    return r.insTpDto
}
// JzReceiverTo Setter
// 家装物流-家装收货人信息,如果为空,则取默认收货信息
func (r *TaobaoOmniorderStoreConsignedRequest) SetJzReceiverTo(jzReceiverTo *JzReceiverDto) error {
    r.jzReceiverTo = jzReceiverTo
    r.Set("jz_receiver_to", jzReceiverTo)
    return nil
}

// JzReceiverTo Getter
func (r TaobaoOmniorderStoreConsignedRequest) GetJzReceiverTo() *JzReceiverDto {
    return r.jzReceiverTo
}
// Tid Setter
// 淘宝交易主订单ID
func (r *TaobaoOmniorderStoreConsignedRequest) SetTid(tid int64) error {
    r.tid = tid
    r.Set("tid", tid)
    return nil
}

// Tid Getter
func (r TaobaoOmniorderStoreConsignedRequest) GetTid() int64 {
    return r.tid
}
// ReportTimestamp Setter
// ISV系统上报时间
func (r *TaobaoOmniorderStoreConsignedRequest) SetReportTimestamp(reportTimestamp int64) error {
    r.reportTimestamp = reportTimestamp
    r.Set("report_timestamp", reportTimestamp)
    return nil
}

// ReportTimestamp Getter
func (r TaobaoOmniorderStoreConsignedRequest) GetReportTimestamp() int64 {
    return r.reportTimestamp
}
// LgTpDto Setter
// 家装物流-物流公司信息
func (r *TaobaoOmniorderStoreConsignedRequest) SetLgTpDto(lgTpDto *TpDto) error {
    r.lgTpDto = lgTpDto
    r.Set("lg_tp_dto", lgTpDto)
    return nil
}

// LgTpDto Getter
func (r TaobaoOmniorderStoreConsignedRequest) GetLgTpDto() *TpDto {
    return r.lgTpDto
}

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
    _traceId   string
    // 卖家联系人地址库ID，可以通过taobao.logistics.address.search接口查询到地址库ID。如果为空，取的卖家的默认取货地址
    _senderId   int64
    // 家装物流-安装收货人信息,如果为空,则取默认收货人信息
    _insReceiverTo   *JzReceiverDTO
    // 子订单列表
    _subOrderList   []StoreConsignedResult
    // 家装物流-发货参数
    _jzTopArgs   *JzTopArgsDTO
    // 家装物流-安装公司信息,需要安装时,才填写
    _insTpDto   *TpDTO
    // 家装物流-家装收货人信息,如果为空,则取默认收货信息
    _jzReceiverTo   *JzReceiverDTO
    // 淘宝交易主订单ID
    _tid   int64
    // ISV系统上报时间
    _reportTimestamp   int64
    // 家装物流-物流公司信息
    _lgTpDto   *TpDTO
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
func (r *TaobaoOmniorderStoreConsignedRequest) SetTraceId(_traceId string) error {
    r._traceId = _traceId
    r.Set("trace_id", _traceId)
    return nil
}

// TraceId Getter
func (r TaobaoOmniorderStoreConsignedRequest) GetTraceId() string {
    return r._traceId
}
// SenderId Setter
// 卖家联系人地址库ID，可以通过taobao.logistics.address.search接口查询到地址库ID。如果为空，取的卖家的默认取货地址
func (r *TaobaoOmniorderStoreConsignedRequest) SetSenderId(_senderId int64) error {
    r._senderId = _senderId
    r.Set("sender_id", _senderId)
    return nil
}

// SenderId Getter
func (r TaobaoOmniorderStoreConsignedRequest) GetSenderId() int64 {
    return r._senderId
}
// InsReceiverTo Setter
// 家装物流-安装收货人信息,如果为空,则取默认收货人信息
func (r *TaobaoOmniorderStoreConsignedRequest) SetInsReceiverTo(_insReceiverTo *JzReceiverDTO) error {
    r._insReceiverTo = _insReceiverTo
    r.Set("ins_receiver_to", _insReceiverTo)
    return nil
}

// InsReceiverTo Getter
func (r TaobaoOmniorderStoreConsignedRequest) GetInsReceiverTo() *JzReceiverDTO {
    return r._insReceiverTo
}
// SubOrderList Setter
// 子订单列表
func (r *TaobaoOmniorderStoreConsignedRequest) SetSubOrderList(_subOrderList []StoreConsignedResult) error {
    r._subOrderList = _subOrderList
    r.Set("sub_order_list", _subOrderList)
    return nil
}

// SubOrderList Getter
func (r TaobaoOmniorderStoreConsignedRequest) GetSubOrderList() []StoreConsignedResult {
    return r._subOrderList
}
// JzTopArgs Setter
// 家装物流-发货参数
func (r *TaobaoOmniorderStoreConsignedRequest) SetJzTopArgs(_jzTopArgs *JzTopArgsDTO) error {
    r._jzTopArgs = _jzTopArgs
    r.Set("jz_top_args", _jzTopArgs)
    return nil
}

// JzTopArgs Getter
func (r TaobaoOmniorderStoreConsignedRequest) GetJzTopArgs() *JzTopArgsDTO {
    return r._jzTopArgs
}
// InsTpDto Setter
// 家装物流-安装公司信息,需要安装时,才填写
func (r *TaobaoOmniorderStoreConsignedRequest) SetInsTpDto(_insTpDto *TpDTO) error {
    r._insTpDto = _insTpDto
    r.Set("ins_tp_dto", _insTpDto)
    return nil
}

// InsTpDto Getter
func (r TaobaoOmniorderStoreConsignedRequest) GetInsTpDto() *TpDTO {
    return r._insTpDto
}
// JzReceiverTo Setter
// 家装物流-家装收货人信息,如果为空,则取默认收货信息
func (r *TaobaoOmniorderStoreConsignedRequest) SetJzReceiverTo(_jzReceiverTo *JzReceiverDTO) error {
    r._jzReceiverTo = _jzReceiverTo
    r.Set("jz_receiver_to", _jzReceiverTo)
    return nil
}

// JzReceiverTo Getter
func (r TaobaoOmniorderStoreConsignedRequest) GetJzReceiverTo() *JzReceiverDTO {
    return r._jzReceiverTo
}
// Tid Setter
// 淘宝交易主订单ID
func (r *TaobaoOmniorderStoreConsignedRequest) SetTid(_tid int64) error {
    r._tid = _tid
    r.Set("tid", _tid)
    return nil
}

// Tid Getter
func (r TaobaoOmniorderStoreConsignedRequest) GetTid() int64 {
    return r._tid
}
// ReportTimestamp Setter
// ISV系统上报时间
func (r *TaobaoOmniorderStoreConsignedRequest) SetReportTimestamp(_reportTimestamp int64) error {
    r._reportTimestamp = _reportTimestamp
    r.Set("report_timestamp", _reportTimestamp)
    return nil
}

// ReportTimestamp Getter
func (r TaobaoOmniorderStoreConsignedRequest) GetReportTimestamp() int64 {
    return r._reportTimestamp
}
// LgTpDto Setter
// 家装物流-物流公司信息
func (r *TaobaoOmniorderStoreConsignedRequest) SetLgTpDto(_lgTpDto *TpDTO) error {
    r._lgTpDto = _lgTpDto
    r.Set("lg_tp_dto", _lgTpDto)
    return nil
}

// LgTpDto Getter
func (r TaobaoOmniorderStoreConsignedRequest) GetLgTpDto() *TpDTO {
    return r._lgTpDto
}

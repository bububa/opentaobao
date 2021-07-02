package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOmniorderStoreConsignedAPIRequest Pos端门店发货 API请求
// taobao.omniorder.store.consigned
//
// ISV Pos端门店发货，通知星盘
type TaobaoOmniorderStoreConsignedAPIRequest struct {
	model.Params
	// 跟踪Id
	_traceId string
	// 卖家联系人地址库ID，可以通过taobao.logistics.address.search接口查询到地址库ID。如果为空，取的卖家的默认取货地址
	_senderId int64
	// 家装物流-安装收货人信息,如果为空,则取默认收货人信息
	_insReceiverTo *JzReceiverDto
	// 子订单列表
	_subOrderList []StoreConsignedResult
	// 家装物流-发货参数
	_jzTopArgs *JzTopArgsDto
	// 家装物流-安装公司信息,需要安装时,才填写
	_insTpDto *TpDto
	// 家装物流-家装收货人信息,如果为空,则取默认收货信息
	_jzReceiverTo *JzReceiverDto
	// 淘宝交易主订单ID
	_tid int64
	// ISV系统上报时间
	_reportTimestamp int64
	// 家装物流-物流公司信息
	_lgTpDto *TpDto
}

// NewTaobaoOmniorderStoreConsignedRequest 初始化TaobaoOmniorderStoreConsignedAPIRequest对象
func NewTaobaoOmniorderStoreConsignedRequest() *TaobaoOmniorderStoreConsignedAPIRequest {
	return &TaobaoOmniorderStoreConsignedAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOmniorderStoreConsignedAPIRequest) GetApiMethodName() string {
	return "taobao.omniorder.store.consigned"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOmniorderStoreConsignedAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetTraceId is TraceId Setter
// 跟踪Id
func (r *TaobaoOmniorderStoreConsignedAPIRequest) SetTraceId(_traceId string) error {
	r._traceId = _traceId
	r.Set("trace_id", _traceId)
	return nil
}

// GetTraceId TraceId Getter
func (r TaobaoOmniorderStoreConsignedAPIRequest) GetTraceId() string {
	return r._traceId
}

// SetSenderId is SenderId Setter
// 卖家联系人地址库ID，可以通过taobao.logistics.address.search接口查询到地址库ID。如果为空，取的卖家的默认取货地址
func (r *TaobaoOmniorderStoreConsignedAPIRequest) SetSenderId(_senderId int64) error {
	r._senderId = _senderId
	r.Set("sender_id", _senderId)
	return nil
}

// GetSenderId SenderId Getter
func (r TaobaoOmniorderStoreConsignedAPIRequest) GetSenderId() int64 {
	return r._senderId
}

// SetInsReceiverTo is InsReceiverTo Setter
// 家装物流-安装收货人信息,如果为空,则取默认收货人信息
func (r *TaobaoOmniorderStoreConsignedAPIRequest) SetInsReceiverTo(_insReceiverTo *JzReceiverDto) error {
	r._insReceiverTo = _insReceiverTo
	r.Set("ins_receiver_to", _insReceiverTo)
	return nil
}

// GetInsReceiverTo InsReceiverTo Getter
func (r TaobaoOmniorderStoreConsignedAPIRequest) GetInsReceiverTo() *JzReceiverDto {
	return r._insReceiverTo
}

// SetSubOrderList is SubOrderList Setter
// 子订单列表
func (r *TaobaoOmniorderStoreConsignedAPIRequest) SetSubOrderList(_subOrderList []StoreConsignedResult) error {
	r._subOrderList = _subOrderList
	r.Set("sub_order_list", _subOrderList)
	return nil
}

// GetSubOrderList SubOrderList Getter
func (r TaobaoOmniorderStoreConsignedAPIRequest) GetSubOrderList() []StoreConsignedResult {
	return r._subOrderList
}

// SetJzTopArgs is JzTopArgs Setter
// 家装物流-发货参数
func (r *TaobaoOmniorderStoreConsignedAPIRequest) SetJzTopArgs(_jzTopArgs *JzTopArgsDto) error {
	r._jzTopArgs = _jzTopArgs
	r.Set("jz_top_args", _jzTopArgs)
	return nil
}

// GetJzTopArgs JzTopArgs Getter
func (r TaobaoOmniorderStoreConsignedAPIRequest) GetJzTopArgs() *JzTopArgsDto {
	return r._jzTopArgs
}

// SetInsTpDto is InsTpDto Setter
// 家装物流-安装公司信息,需要安装时,才填写
func (r *TaobaoOmniorderStoreConsignedAPIRequest) SetInsTpDto(_insTpDto *TpDto) error {
	r._insTpDto = _insTpDto
	r.Set("ins_tp_dto", _insTpDto)
	return nil
}

// GetInsTpDto InsTpDto Getter
func (r TaobaoOmniorderStoreConsignedAPIRequest) GetInsTpDto() *TpDto {
	return r._insTpDto
}

// SetJzReceiverTo is JzReceiverTo Setter
// 家装物流-家装收货人信息,如果为空,则取默认收货信息
func (r *TaobaoOmniorderStoreConsignedAPIRequest) SetJzReceiverTo(_jzReceiverTo *JzReceiverDto) error {
	r._jzReceiverTo = _jzReceiverTo
	r.Set("jz_receiver_to", _jzReceiverTo)
	return nil
}

// GetJzReceiverTo JzReceiverTo Getter
func (r TaobaoOmniorderStoreConsignedAPIRequest) GetJzReceiverTo() *JzReceiverDto {
	return r._jzReceiverTo
}

// SetTid is Tid Setter
// 淘宝交易主订单ID
func (r *TaobaoOmniorderStoreConsignedAPIRequest) SetTid(_tid int64) error {
	r._tid = _tid
	r.Set("tid", _tid)
	return nil
}

// GetTid Tid Getter
func (r TaobaoOmniorderStoreConsignedAPIRequest) GetTid() int64 {
	return r._tid
}

// SetReportTimestamp is ReportTimestamp Setter
// ISV系统上报时间
func (r *TaobaoOmniorderStoreConsignedAPIRequest) SetReportTimestamp(_reportTimestamp int64) error {
	r._reportTimestamp = _reportTimestamp
	r.Set("report_timestamp", _reportTimestamp)
	return nil
}

// GetReportTimestamp ReportTimestamp Getter
func (r TaobaoOmniorderStoreConsignedAPIRequest) GetReportTimestamp() int64 {
	return r._reportTimestamp
}

// SetLgTpDto is LgTpDto Setter
// 家装物流-物流公司信息
func (r *TaobaoOmniorderStoreConsignedAPIRequest) SetLgTpDto(_lgTpDto *TpDto) error {
	r._lgTpDto = _lgTpDto
	r.Set("lg_tp_dto", _lgTpDto)
	return nil
}

// GetLgTpDto LgTpDto Getter
func (r TaobaoOmniorderStoreConsignedAPIRequest) GetLgTpDto() *TpDto {
	return r._lgTpDto
}

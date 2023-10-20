package tblogistics

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWlbOrderJzConsignAPIRequest 家装发货接口 API请求
// taobao.wlb.order.jz.consign
//
// 家装类订单使用该接口发货
type TaobaoWlbOrderJzConsignAPIRequest struct {
	model.Params
	// 交易号
	_tid int64
	// 卖家联系人地址库ID，可以通过taobao.logistics.address.search接口查询到地址库ID。如果为空，取的卖家的默认取货地址
	_senderId int64
	// 物流公司信息
	_lgTpDto *TPDto
	// 安装公司信息,需要安装时,才填写
	_insTpDto *TPDto
	// 家装收货人信息,如果为空,则取默认收货信息
	_jzReceiverTo *JzReceiverTO
	// 安装收货人信息,如果为空,则取默认收货人信息
	_insReceiverTo *JzReceiverTO
	// 发货参数
	_jzTopArgs *JzTopArgs
}

// NewTaobaoWlbOrderJzConsignRequest 初始化TaobaoWlbOrderJzConsignAPIRequest对象
func NewTaobaoWlbOrderJzConsignRequest() *TaobaoWlbOrderJzConsignAPIRequest {
	return &TaobaoWlbOrderJzConsignAPIRequest{
		Params: model.NewParams(7),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoWlbOrderJzConsignAPIRequest) Reset() {
	r._tid = 0
	r._senderId = 0
	r._lgTpDto = nil
	r._insTpDto = nil
	r._jzReceiverTo = nil
	r._insReceiverTo = nil
	r._jzTopArgs = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoWlbOrderJzConsignAPIRequest) GetApiMethodName() string {
	return "taobao.wlb.order.jz.consign"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoWlbOrderJzConsignAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoWlbOrderJzConsignAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTid is Tid Setter
// 交易号
func (r *TaobaoWlbOrderJzConsignAPIRequest) SetTid(_tid int64) error {
	r._tid = _tid
	r.Set("tid", _tid)
	return nil
}

// GetTid Tid Getter
func (r TaobaoWlbOrderJzConsignAPIRequest) GetTid() int64 {
	return r._tid
}

// SetSenderId is SenderId Setter
// 卖家联系人地址库ID，可以通过taobao.logistics.address.search接口查询到地址库ID。如果为空，取的卖家的默认取货地址
func (r *TaobaoWlbOrderJzConsignAPIRequest) SetSenderId(_senderId int64) error {
	r._senderId = _senderId
	r.Set("sender_id", _senderId)
	return nil
}

// GetSenderId SenderId Getter
func (r TaobaoWlbOrderJzConsignAPIRequest) GetSenderId() int64 {
	return r._senderId
}

// SetLgTpDto is LgTpDto Setter
// 物流公司信息
func (r *TaobaoWlbOrderJzConsignAPIRequest) SetLgTpDto(_lgTpDto *TPDto) error {
	r._lgTpDto = _lgTpDto
	r.Set("lg_tp_dto", _lgTpDto)
	return nil
}

// GetLgTpDto LgTpDto Getter
func (r TaobaoWlbOrderJzConsignAPIRequest) GetLgTpDto() *TPDto {
	return r._lgTpDto
}

// SetInsTpDto is InsTpDto Setter
// 安装公司信息,需要安装时,才填写
func (r *TaobaoWlbOrderJzConsignAPIRequest) SetInsTpDto(_insTpDto *TPDto) error {
	r._insTpDto = _insTpDto
	r.Set("ins_tp_dto", _insTpDto)
	return nil
}

// GetInsTpDto InsTpDto Getter
func (r TaobaoWlbOrderJzConsignAPIRequest) GetInsTpDto() *TPDto {
	return r._insTpDto
}

// SetJzReceiverTo is JzReceiverTo Setter
// 家装收货人信息,如果为空,则取默认收货信息
func (r *TaobaoWlbOrderJzConsignAPIRequest) SetJzReceiverTo(_jzReceiverTo *JzReceiverTO) error {
	r._jzReceiverTo = _jzReceiverTo
	r.Set("jz_receiver_to", _jzReceiverTo)
	return nil
}

// GetJzReceiverTo JzReceiverTo Getter
func (r TaobaoWlbOrderJzConsignAPIRequest) GetJzReceiverTo() *JzReceiverTO {
	return r._jzReceiverTo
}

// SetInsReceiverTo is InsReceiverTo Setter
// 安装收货人信息,如果为空,则取默认收货人信息
func (r *TaobaoWlbOrderJzConsignAPIRequest) SetInsReceiverTo(_insReceiverTo *JzReceiverTO) error {
	r._insReceiverTo = _insReceiverTo
	r.Set("ins_receiver_to", _insReceiverTo)
	return nil
}

// GetInsReceiverTo InsReceiverTo Getter
func (r TaobaoWlbOrderJzConsignAPIRequest) GetInsReceiverTo() *JzReceiverTO {
	return r._insReceiverTo
}

// SetJzTopArgs is JzTopArgs Setter
// 发货参数
func (r *TaobaoWlbOrderJzConsignAPIRequest) SetJzTopArgs(_jzTopArgs *JzTopArgs) error {
	r._jzTopArgs = _jzTopArgs
	r.Set("jz_top_args", _jzTopArgs)
	return nil
}

// GetJzTopArgs JzTopArgs Getter
func (r TaobaoWlbOrderJzConsignAPIRequest) GetJzTopArgs() *JzTopArgs {
	return r._jzTopArgs
}

var poolTaobaoWlbOrderJzConsignAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoWlbOrderJzConsignRequest()
	},
}

// GetTaobaoWlbOrderJzConsignRequest 从 sync.Pool 获取 TaobaoWlbOrderJzConsignAPIRequest
func GetTaobaoWlbOrderJzConsignAPIRequest() *TaobaoWlbOrderJzConsignAPIRequest {
	return poolTaobaoWlbOrderJzConsignAPIRequest.Get().(*TaobaoWlbOrderJzConsignAPIRequest)
}

// ReleaseTaobaoWlbOrderJzConsignAPIRequest 将 TaobaoWlbOrderJzConsignAPIRequest 放入 sync.Pool
func ReleaseTaobaoWlbOrderJzConsignAPIRequest(v *TaobaoWlbOrderJzConsignAPIRequest) {
	v.Reset()
	poolTaobaoWlbOrderJzConsignAPIRequest.Put(v)
}

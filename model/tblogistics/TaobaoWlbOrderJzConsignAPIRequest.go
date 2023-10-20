package tblogistics

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaowlborderjzconsignAPIRequest 家装发货接口 API请求
// taobao.wlb.order.jz.consign
//
// 家装类订单使用该接口发货
type TaobaowlborderjzconsignAPIRequest struct {
	model.Params
	// 交易号
	_tid int64
	// 卖家联系人地址库ID，可以通过taobao.logistics.address.search接口查询到地址库ID。如果为空，取的卖家的默认取货地址
	_senderId int64
	// 物流公司信息
	_lgTpDto *Tpdto
	// 安装公司信息,需要安装时,才填写
	_insTpDto *Tpdto
	// 家装收货人信息,如果为空,则取默认收货信息
	_jzReceiverTo *JzReceiverTo
	// 安装收货人信息,如果为空,则取默认收货人信息
	_insReceiverTo *JzReceiverTo
	// 发货参数
	_jzTopArgs *JzTopArgs
}

// NewTaobaowlborderjzconsignRequest 初始化TaobaowlborderjzconsignAPIRequest对象
func NewTaobaowlborderjzconsignRequest() *TaobaowlborderjzconsignAPIRequest {
	return &TaobaowlborderjzconsignAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaowlborderjzconsignAPIRequest) GetApiMethodName() string {
	return "taobao.wlb.order.jz.consign"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaowlborderjzconsignAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaowlborderjzconsignAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTid is Tid Setter
// 交易号
func (r *TaobaowlborderjzconsignAPIRequest) SetTid(_tid int64) error {
	r._tid = _tid
	r.Set("tid", _tid)
	return nil
}

// GetTid Tid Getter
func (r TaobaowlborderjzconsignAPIRequest) GetTid() int64 {
	return r._tid
}

// SetSenderId is SenderId Setter
// 卖家联系人地址库ID，可以通过taobao.logistics.address.search接口查询到地址库ID。如果为空，取的卖家的默认取货地址
func (r *TaobaowlborderjzconsignAPIRequest) SetSenderId(_senderId int64) error {
	r._senderId = _senderId
	r.Set("sender_id", _senderId)
	return nil
}

// GetSenderId SenderId Getter
func (r TaobaowlborderjzconsignAPIRequest) GetSenderId() int64 {
	return r._senderId
}

// SetLgTpDto is LgTpDto Setter
// 物流公司信息
func (r *TaobaowlborderjzconsignAPIRequest) SetLgTpDto(_lgTpDto *Tpdto) error {
	r._lgTpDto = _lgTpDto
	r.Set("lg_tp_dto", _lgTpDto)
	return nil
}

// GetLgTpDto LgTpDto Getter
func (r TaobaowlborderjzconsignAPIRequest) GetLgTpDto() *Tpdto {
	return r._lgTpDto
}

// SetInsTpDto is InsTpDto Setter
// 安装公司信息,需要安装时,才填写
func (r *TaobaowlborderjzconsignAPIRequest) SetInsTpDto(_insTpDto *Tpdto) error {
	r._insTpDto = _insTpDto
	r.Set("ins_tp_dto", _insTpDto)
	return nil
}

// GetInsTpDto InsTpDto Getter
func (r TaobaowlborderjzconsignAPIRequest) GetInsTpDto() *Tpdto {
	return r._insTpDto
}

// SetJzReceiverTo is JzReceiverTo Setter
// 家装收货人信息,如果为空,则取默认收货信息
func (r *TaobaowlborderjzconsignAPIRequest) SetJzReceiverTo(_jzReceiverTo *JzReceiverTo) error {
	r._jzReceiverTo = _jzReceiverTo
	r.Set("jz_receiver_to", _jzReceiverTo)
	return nil
}

// GetJzReceiverTo JzReceiverTo Getter
func (r TaobaowlborderjzconsignAPIRequest) GetJzReceiverTo() *JzReceiverTo {
	return r._jzReceiverTo
}

// SetInsReceiverTo is InsReceiverTo Setter
// 安装收货人信息,如果为空,则取默认收货人信息
func (r *TaobaowlborderjzconsignAPIRequest) SetInsReceiverTo(_insReceiverTo *JzReceiverTo) error {
	r._insReceiverTo = _insReceiverTo
	r.Set("ins_receiver_to", _insReceiverTo)
	return nil
}

// GetInsReceiverTo InsReceiverTo Getter
func (r TaobaowlborderjzconsignAPIRequest) GetInsReceiverTo() *JzReceiverTo {
	return r._insReceiverTo
}

// SetJzTopArgs is JzTopArgs Setter
// 发货参数
func (r *TaobaowlborderjzconsignAPIRequest) SetJzTopArgs(_jzTopArgs *JzTopArgs) error {
	r._jzTopArgs = _jzTopArgs
	r.Set("jz_top_args", _jzTopArgs)
	return nil
}

// GetJzTopArgs JzTopArgs Getter
func (r TaobaowlborderjzconsignAPIRequest) GetJzTopArgs() *JzTopArgs {
	return r._jzTopArgs
}

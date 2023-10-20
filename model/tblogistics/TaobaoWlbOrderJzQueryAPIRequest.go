package tblogistics

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWlbOrderJzQueryAPIRequest 家装业务查询物流公司api API请求
// taobao.wlb.order.jz.query
//
// 家装业务查询物流公司api
type TaobaoWlbOrderJzQueryAPIRequest struct {
	model.Params
	// 交易id
	_tid int64
	// 卖家联系人地址库ID，可以通过taobao.logistics.address.search接口查询到地址库ID。如果为空，取的卖家的默认取货地址
	_senderId int64
	// 家装收货人信息
	_jzReceiverTo *JzReceiverTO
	// 家装安装服务收货人信息
	_insJzReceiverTO *JzReceiverTO
}

// NewTaobaoWlbOrderJzQueryRequest 初始化TaobaoWlbOrderJzQueryAPIRequest对象
func NewTaobaoWlbOrderJzQueryRequest() *TaobaoWlbOrderJzQueryAPIRequest {
	return &TaobaoWlbOrderJzQueryAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoWlbOrderJzQueryAPIRequest) Reset() {
	r._tid = 0
	r._senderId = 0
	r._jzReceiverTo = nil
	r._insJzReceiverTO = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoWlbOrderJzQueryAPIRequest) GetApiMethodName() string {
	return "taobao.wlb.order.jz.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoWlbOrderJzQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoWlbOrderJzQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTid is Tid Setter
// 交易id
func (r *TaobaoWlbOrderJzQueryAPIRequest) SetTid(_tid int64) error {
	r._tid = _tid
	r.Set("tid", _tid)
	return nil
}

// GetTid Tid Getter
func (r TaobaoWlbOrderJzQueryAPIRequest) GetTid() int64 {
	return r._tid
}

// SetSenderId is SenderId Setter
// 卖家联系人地址库ID，可以通过taobao.logistics.address.search接口查询到地址库ID。如果为空，取的卖家的默认取货地址
func (r *TaobaoWlbOrderJzQueryAPIRequest) SetSenderId(_senderId int64) error {
	r._senderId = _senderId
	r.Set("sender_id", _senderId)
	return nil
}

// GetSenderId SenderId Getter
func (r TaobaoWlbOrderJzQueryAPIRequest) GetSenderId() int64 {
	return r._senderId
}

// SetJzReceiverTo is JzReceiverTo Setter
// 家装收货人信息
func (r *TaobaoWlbOrderJzQueryAPIRequest) SetJzReceiverTo(_jzReceiverTo *JzReceiverTO) error {
	r._jzReceiverTo = _jzReceiverTo
	r.Set("jz_receiver_to", _jzReceiverTo)
	return nil
}

// GetJzReceiverTo JzReceiverTo Getter
func (r TaobaoWlbOrderJzQueryAPIRequest) GetJzReceiverTo() *JzReceiverTO {
	return r._jzReceiverTo
}

// SetInsJzReceiverTO is InsJzReceiverTO Setter
// 家装安装服务收货人信息
func (r *TaobaoWlbOrderJzQueryAPIRequest) SetInsJzReceiverTO(_insJzReceiverTO *JzReceiverTO) error {
	r._insJzReceiverTO = _insJzReceiverTO
	r.Set("ins_jz_receiver_t_o", _insJzReceiverTO)
	return nil
}

// GetInsJzReceiverTO InsJzReceiverTO Getter
func (r TaobaoWlbOrderJzQueryAPIRequest) GetInsJzReceiverTO() *JzReceiverTO {
	return r._insJzReceiverTO
}

var poolTaobaoWlbOrderJzQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoWlbOrderJzQueryRequest()
	},
}

// GetTaobaoWlbOrderJzQueryRequest 从 sync.Pool 获取 TaobaoWlbOrderJzQueryAPIRequest
func GetTaobaoWlbOrderJzQueryAPIRequest() *TaobaoWlbOrderJzQueryAPIRequest {
	return poolTaobaoWlbOrderJzQueryAPIRequest.Get().(*TaobaoWlbOrderJzQueryAPIRequest)
}

// ReleaseTaobaoWlbOrderJzQueryAPIRequest 将 TaobaoWlbOrderJzQueryAPIRequest 放入 sync.Pool
func ReleaseTaobaoWlbOrderJzQueryAPIRequest(v *TaobaoWlbOrderJzQueryAPIRequest) {
	v.Reset()
	poolTaobaoWlbOrderJzQueryAPIRequest.Put(v)
}

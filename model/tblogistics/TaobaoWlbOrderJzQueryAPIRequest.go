package tblogistics

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaowlborderjzqueryAPIRequest 家装业务查询物流公司api API请求
// taobao.wlb.order.jz.query
//
// 家装业务查询物流公司api
type TaobaowlborderjzqueryAPIRequest struct {
	model.Params
	// 交易id
	_tid int64
	// 卖家联系人地址库ID，可以通过taobao.logistics.address.search接口查询到地址库ID。如果为空，取的卖家的默认取货地址
	_senderId int64
	// 家装收货人信息
	_jzReceiverTo *JzReceiverTo
	// 家装安装服务收货人信息
	_insJzReceiverTO *JzReceiverTo
}

// NewTaobaowlborderjzqueryRequest 初始化TaobaowlborderjzqueryAPIRequest对象
func NewTaobaowlborderjzqueryRequest() *TaobaowlborderjzqueryAPIRequest {
	return &TaobaowlborderjzqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaowlborderjzqueryAPIRequest) GetApiMethodName() string {
	return "taobao.wlb.order.jz.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaowlborderjzqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaowlborderjzqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTid is Tid Setter
// 交易id
func (r *TaobaowlborderjzqueryAPIRequest) SetTid(_tid int64) error {
	r._tid = _tid
	r.Set("tid", _tid)
	return nil
}

// GetTid Tid Getter
func (r TaobaowlborderjzqueryAPIRequest) GetTid() int64 {
	return r._tid
}

// SetSenderId is SenderId Setter
// 卖家联系人地址库ID，可以通过taobao.logistics.address.search接口查询到地址库ID。如果为空，取的卖家的默认取货地址
func (r *TaobaowlborderjzqueryAPIRequest) SetSenderId(_senderId int64) error {
	r._senderId = _senderId
	r.Set("sender_id", _senderId)
	return nil
}

// GetSenderId SenderId Getter
func (r TaobaowlborderjzqueryAPIRequest) GetSenderId() int64 {
	return r._senderId
}

// SetJzReceiverTo is JzReceiverTo Setter
// 家装收货人信息
func (r *TaobaowlborderjzqueryAPIRequest) SetJzReceiverTo(_jzReceiverTo *JzReceiverTo) error {
	r._jzReceiverTo = _jzReceiverTo
	r.Set("jz_receiver_to", _jzReceiverTo)
	return nil
}

// GetJzReceiverTo JzReceiverTo Getter
func (r TaobaowlborderjzqueryAPIRequest) GetJzReceiverTo() *JzReceiverTo {
	return r._jzReceiverTo
}

// SetInsJzReceiverTO is InsJzReceiverTO Setter
// 家装安装服务收货人信息
func (r *TaobaowlborderjzqueryAPIRequest) SetInsJzReceiverTO(_insJzReceiverTO *JzReceiverTo) error {
	r._insJzReceiverTO = _insJzReceiverTO
	r.Set("ins_jz_receiver_t_o", _insJzReceiverTO)
	return nil
}

// GetInsJzReceiverTO InsJzReceiverTO Getter
func (r TaobaowlborderjzqueryAPIRequest) GetInsJzReceiverTO() *JzReceiverTo {
	return r._insJzReceiverTO
}

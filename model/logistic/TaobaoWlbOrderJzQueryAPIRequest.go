package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoWlbOrderJzQueryAPIRequest
家装业务查询物流公司api API请求
taobao.wlb.order.jz.query

家装业务查询物流公司api */
type TaobaoWlbOrderJzQueryAPIRequest struct {
	model.Params
	// 交易id
	_tid int64
	// 家装收货人信息
	_jzReceiverTo *JzReceiverTo
	// 卖家联系人地址库ID，可以通过taobao.logistics.address.search接口查询到地址库ID。如果为空，取的卖家的默认取货地址
	_senderId int64
	// 家装安装服务收货人信息
	_insJzReceiverTO *JzReceiverTo
}

// New

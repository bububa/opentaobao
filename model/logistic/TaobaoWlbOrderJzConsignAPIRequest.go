package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoWlbOrderJzConsignAPIRequest
家装发货接口 API请求
taobao.wlb.order.jz.consign

家装类订单使用该接口发货 */
type TaobaoWlbOrderJzConsignAPIRequest struct {
	model.Params
	// 交易号
	_tid int64
	// 卖家联系人地址库ID，可以通过taobao.logistics.address.search接口查询到地址库ID。如果为空，取的卖家的默认取货地址
	_senderId int64
	// 家装收货人信息,如果为空,则取默认收货信息
	_jzReceiverTo *JzReceiverTo
	// 发货参数
	_jzTopArgs *JzTopArgs
	// 物流公司信息
	_lgTpDto *Tpdto
	// 安装公司信息,需要安装时,才填写
	_insTpDto *Tpdto
	// 安装收货人信息,如果为空,则取默认收货人信息
	_insReceiverTo *JzReceiverTo
}

// New

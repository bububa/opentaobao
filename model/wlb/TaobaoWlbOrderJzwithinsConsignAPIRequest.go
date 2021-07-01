package wlb

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoWlbOrderJzwithinsConsignAPIRequest
家装发货接口 API请求
taobao.wlb.order.jzwithins.consign

为支持家装类目的商家，对绑定家装物流服务的订单可以在商家的ERP中发货、批量发货，因此开发带安装服务商的发货接口 */
type TaobaoWlbOrderJzwithinsConsignAPIRequest struct {
	model.Params
	// 淘宝交易订单号
	_tid int64
	// 物流服务商信息
	_tmsPartner *JzPartnerNew
	// 物流服务商信息
	_insPartner *JzPartnerNew
	// 家装物流发货参数
	_jzConsignArgs *JzConsignArgsNew
}

// New

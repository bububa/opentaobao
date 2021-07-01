package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoLogisticsConsignOrderCreateandsendAPIRequest
创建订单并发货 API请求
taobao.logistics.consign.order.createandsend

创建物流订单，并发货。 */
type TaobaoLogisticsConsignOrderCreateandsendAPIRequest struct {
	model.Params
	// 用户ID
	_userId int64
	// 订单来源，值选择：30
	_orderSource int64
	// 订单类型，值固定选择：30
	_orderType int64
	// 物流订单物流类型，值固定选择：2
	_logisType int64
	// 物流公司ID
	_companyId int64
	// 交易流水号，淘外订单号或者商家内部交易流水号
	_tradeId int64
	// 运单号
	_mailNo string
	// 费用承担方式 1买家承担运费 2卖家承担运费
	_shipping string
	// 发件人名称
	_sName string
	// 发件人区域ID
	_sAreaId int64
	// 发件人街道地址
	_sAddress string
	// 发件人出编
	_sZipCode string
	// 手机号码
	_sMobilePhone string
	// 电话号码
	_sTelephone string
	// 省
	_sProvName string
	// 市
	_sCityName string
	// 区
	_sDistName string
	// 收件人名称
	_rName string
	// 收件人区域ID
	_rAreaId int64
	// 收件人街道地址
	_rAddress string
	// 收件人邮编
	_rZipCode string
	// 手机号码
	_rMobilePhone string
	// 电话号码
	_rTelephone string
	// 省
	_rProvName string
	// 市
	_rCityName string
	// 区
	_rDistName string
	// 物品的json数据。
	_itemJsonString string
}

// New

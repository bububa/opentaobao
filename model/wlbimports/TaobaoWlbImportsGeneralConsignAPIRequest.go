package wlbimports

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoWlbImportsGeneralConsignAPIRequest
一般进口发货 API请求
taobao.wlb.imports.general.consign

将订单信息发送到菜鸟海外转运仓；
业务规则：
1）交易订单为待发货状态。
2）单笔订单多个商品，交易金额不能大于1000人民币。 */
type TaobaoWlbImportsGeneralConsignAPIRequest struct {
	model.Params
	// 交易订单id
	_tradeOrderId int64
	// 物流资源ID
	_resourceId int64
	// 仓库编码
	_storeCode string
	// 第1段物流承运商
	_firstLogistics string
	// 第1段物流运单号
	_firstWaybillno string
	// 卖家发货地址库ID；不填的话取默认发货地址；如果填写的senderId对应多个地址，取第一个
	_senderId int64
	// 卖家退货地址库ID；不填写的话取默认退货地址；如果填写的cancelId对应多个地址，取第一个
	_cancelId int64
	// 增值服务编码.多个以逗号分隔
	_vasCode string
}

// New

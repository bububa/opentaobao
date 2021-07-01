package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOmniorderStoreConsignedAPIRequest
Pos端门店发货 API请求
taobao.omniorder.store.consigned

ISV Pos端门店发货，通知星盘 */
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

// New

package wlb

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoWlbOrderPageGetAPIRequest
分页查询物流宝订单 API请求
taobao.wlb.order.page.get

分页查询物流宝订单 */
type TaobaoWlbOrderPageGetAPIRequest struct {
	model.Params
	// 每页多少条
	_pageSize int64
	// 分页的第几页
	_pageNo int64
	// TMS拒签：-100 接收方拒签：-200
	_orderStatus int64
	// 物流订单编号
	_orderCode string
	// 订单类型: （1）NORMAL_OUT ：正常出库 （2）NORMAL_IN：正常入库 （3）RETURN_IN：退货入库 （4）EXCHANGE_OUT：换货出库
	_orderType string
	// 订单子类型： （1）OTHER： 其他 （2）TAOBAO_TRADE： 淘宝交易 （3）OTHER_TRADE：其他交易 （4）ALLCOATE： 调拨 （5）CHECK: 盘点单 （6）PURCHASE: 采购单
	_orderSubType string
	// 查询截止时间
	_endTime string
	// 查询开始时间
	_startTime string
}

// New

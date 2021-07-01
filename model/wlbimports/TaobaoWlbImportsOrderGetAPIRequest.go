package wlbimports

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoWlbImportsOrderGetAPIRequest
物流订单获取 API请求
taobao.wlb.imports.order.get

一般进口物流订单获取 */
type TaobaoWlbImportsOrderGetAPIRequest struct {
	model.Params
	// 交易订单号
	_tradeId int64
	// 交易订单开始创建时间
	_gmtCreateStart string
	// 交易订单结束创建时间
	_gmtCreateEnd string
	// 物流订单状态编码。以下依（物流订单状态编码，描述）的形式列举出来：(TIN_CONSING,发货中),(SENT_WAIT_COMPANY_MAKE_SURE,待仓库收货),(ORDER_CANCELED,订单关闭),(COMPANY_MAKE_SURE,海外仓已揽收),(REJECTED_RECEIVED_BY_COMPANY,海外仓拒收),(ORDER_REFUNDING,退货中),(ORDER_REFUND_BY_COMPANY,订单已退货),(EXCEPTION_IN_COMPANY,海外仓内异常),(FAILED_PAID_SHIPPING_FEE,支付失败),(PAID_SHIPPING_FEE,待仓库发货),(COMPANY_CONSIGN_CONFIRM,海外仓已发货),(WAIT_CUSTOMS_MAKE_SURE,清关已收货),(EXCEPTION_IN_CUSTOMS,清关异常),(CUSTOMSING,清关中),(COMPANY_DELIVERY,国内配送)。
	_statusCode string
	// 页码。取值范围:大于零的整数; 默认值:1
	_pageNo int64
	// 每页条数。取值范围:大于0小于等于100的整数; 默认值:40; 最小值：1；最大值:20
	_pageSize int64
}

// New

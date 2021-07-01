package jst

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOcTradetraceAlertsGetAPIRequest
异常订单信息获取 API请求
taobao.oc.tradetrace.alerts.get

提供订单预警模块的数据查询接口 */
type TaobaoOcTradetraceAlertsGetAPIRequest struct {
	model.Params
	// 异常类型代码：发货回写淘宝异常=1，商家系统漏单提醒=2，发货超时提醒=3，物流寄送超时=4，买家签收超时=5，物流中转异常=6，退货寄回异常=7，订单追回提醒=8"。
	_abnormalType int64
	// 返回数据的页码
	_pageNo int64
	// 返回数据每页包含的记录数目
	_pageSize int64
}

// New

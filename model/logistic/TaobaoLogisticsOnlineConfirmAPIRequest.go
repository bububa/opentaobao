package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoLogisticsOnlineConfirmAPIRequest
确认发货通知接口 API请求
taobao.logistics.online.confirm

<br><font color='red'>仅在使用taobao.logistics.online.send 发货时未输入运单号的情况下，需要使用该接口补充填写运单号，来确认发货。<br>
确认发货的目的是让交易流程继续走下去，确认发货后交易状态会由【买家已付款】变为【卖家已发货】。</font> */
type TaobaoLogisticsOnlineConfirmAPIRequest struct {
	model.Params
	// 淘宝交易ID
	_tid int64
	// 拆单子订单列表，对应的数据是：子订单号的列表。可以不传，但是如果传了则必须符合传递的规则。子订单必须是操作的物流订单的子订单的真子集
	_subTid []int64
	// 表明是否是拆单，默认值0，1表示拆单
	_isSplit int64
	// 运单号.具体一个物流公司的真实运单号码。淘宝官方物流会校验，请谨慎传入；
	_outSid string
	// 商家的IP地址
	_sellerIp string
}

// New

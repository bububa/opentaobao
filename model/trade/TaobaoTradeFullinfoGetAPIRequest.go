package trade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTradeFullinfoGetAPIRequest
获取单笔交易的详细信息 API请求
taobao.trade.fullinfo.get

获取单笔交易的详细信息
<br/>1. 只有单笔订单的情况下Trade数据结构中才包含商品相关的信息
<br/>2. 获取到的Order中的payment字段在单笔子订单时包含物流费用，多笔子订单时不包含物流费用
<br/>3. 获取红包优惠金额可以使用字段 coupon_fee
<br/>4. 请按需获取字段，减少TOP系统的压力
<br/><strong><a href="https://console.open.taobao.com/dingWeb.htm?from=tradeapi" target="_blank">点击查看更多交易API说明</a></strong> */
type TaobaoTradeFullinfoGetAPIRequest struct {
	model.Params
	// 需要返回的字段列表，多个字段用半角逗号分隔，可选值为返回示例中能看到的所有字段。
	_fields string
	// 交易编号
	_tid int64
	// appkey未对接oaid加密，则忽略该字段。对接oaid加密情况下，（收货人+手机号+座机+收货地址+create）5个字段组合成oaid，原始订单上座机为空也满足条件。传true，代表必须返回oaid，生成不了就报isv.oaid-field-miss错误；默认或者传false，满足生成条件则返回oaid，否则为空
	_includeOaid string
}

// New

package xhotelonlineorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoXhotelOrderAlipayfaceSettleAPIRequest
信用住订单结账接口 API请求
taobao.xhotel.order.alipayface.settle

用于离店付订单在客人离店后，发起结账以及扣款等后续动作 */
type TaobaoXhotelOrderAlipayfaceSettleAPIRequest struct {
	model.Params
	// 备注
	_memo string
	// 商家订单号
	_outId string
	// 入住房间号
	_roomNo string
	// 杂费总额(不能为负数)
	_otherFee int64
	// 房费总额(必须大于0)。结账时请按订单原价发起结账卖家优惠由飞猪平台发起扣减
	_totalRoomFee int64
	// 每日房价,json格式(包含日期，价格，税费，低价加价前费用等),如果房价和在阿里旅行下单时发生了变化，必须设置该字段.用于更新阿里旅行端的房价信息,涉及到对用户的优惠信息处理等环节(多间房的时候dailyPriceInfo留空)
	_dailyPriceInfo string
	// 实际离店日期，用于校验总房费收取
	_checkOut string
	// 淘宝订单id,必须填写
	_tid int64
	// 杂费明细,如果otherFee>0则该字段必须设置,并和杂费金额相吻合
	_otherFeeDetail string
	// 单间房明细
	_roomSettleInfoList []RoomSettleInfo
	// 此金额是否包含担保金 0：默认值无意义；1：包含；2：不包含
	_containGuarantee int64
	// 结账变价标识，0未变价，1变价
	_priceChange int64
	// 币种标识，默认人民币
	_currencyCode string
	// 汇率
	_currencyRate string
	// 税和服务费
	_taxAndFee int64
	// 应收金额,大于0时，直接按照此金额扣款，忽略房费和杂费金额。该字段仅适用于自助入住订单场景
	_amount int64
	// 酒店外部编码
	_hotelCode string
}

// New

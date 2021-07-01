package xhoteloffline

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoXhotelOrderOfflineSettlePutAPIRequest
线下信用住结账专用接口 API请求
taobao.xhotel.order.offline.settle.put

线下信用住结账专用接口 */
type TaobaoXhotelOrderOfflineSettlePutAPIRequest struct {
	model.Params
	// 淘宝订单id,必须填写
	_tid int64
	// 房费总额(必须大于0)
	_totalRoomFee int64
	// 杂费总额(不能为负数)
	_otherFee int64
	// 杂费明细,如果otherFee>0则该字段必须设置,并和杂费金额相吻合
	_otherFeeDetail string
	// 商家订单号
	_outId string
	// 入住房间号
	_roomNo string
	// 每日房价,json格式,如果房价和在阿里旅行下单时发生了变化，必须设置该字段.用于更新阿里旅行端的房价信息,涉及到对用户的优惠信息处理等环节(多间房的时候dailyPriceInfo留空)
	_dailyPriceInfo string
	// 实际离店日期，用于校验总房费收取
	_checkOut string
	// 备注
	_memo string
	// 房间明细列表
	_roomSettleInfoList []RoomSettleInfo
	// 此金额是否包含担保金 0：默认值无意义；1：包含；2：不包含（多间房结账必须传入）
	_containGuarantee int64
	// 结账请求流水号
	_outUuid string
	// 请求结果通知地址（暂时无效，无需传入）
	_notifyUrl string
	// 应收金额,大于0时，直接按照此金额扣款，忽略房费和杂费金额(单位分)
	_amount int64
	// 商家酒店编码
	_hotelCode string
	// 系统商标识
	_vendor string
}

// New

package xhotel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoXhotelOrderOfficialQualificationGetAPIRequest
官网信用住用户资质校验 API请求
taobao.xhotel.order.official.qualification.get

官网信用住在下单前对用户进行资质校验，资质校验通过才能进行信用支付 */
type TaobaoXhotelOrderOfficialQualificationGetAPIRequest struct {
	model.Params
	// 卖家接收阿里旅行订单状态变更的服务地址（需要实现阿里旅行提供的服务通知规范）
	_notifyUrl string
	// 阿里旅行支付（下单）结束后跳转卖家的页面地址（必须）
	_returnUrl string
	// 扩展字段，json串，用户后续的营销、统计、定制等需求，目前已有key列表：      is_new_user：是否是卖家新用户，1-是，0或者key为null，表示不是      is_first_stay：是否是卖家首住，1-是，0或者key为null，表示不是     （已有列表必须传递）
	_extendAttrs string
	// 用户手机号(可选)
	_mobileNo string
	// 商家在淘宝给分配的渠道名（建议填充较好）
	_vendor string
	// 证件类型, 默认0:身份证; 1: 护照; 2:警官证; 3:士兵证; 4: 回乡证。目前只支持身份证
	_idType int64
	// 加密方式, 默认0: 不加密, 信息会通过淘宝开放平台传输, 阿里旅行可以获取到具体信息;      * 目前只支持不加密
	_encryptType int64
	// 入住人姓名（必选）
	_guestName string
	// 用户支付宝唯一识别码(可选)
	_alipayAccount string
	// 外部会员账号（必选）
	_outMemberAccount string
	// 身份证号，必选
	_idNumber string
	// 每日房价,json格式 ，如果是多间房，则是每日多间房总房价(可选)      * eg:{"day":"2015-08-12","price":48800},      {"day":"2015-08-13","price":48800}
	_dailyPriceInfo string
	// 客人离店日期, 最多支持9间夜
	_checkOut string
	// 客人入住日期
	_checkIn string
	// 外部请求序列表号\流水号，单次请求的唯一标识(必须)
	_outUUID string
	// 总的收费金额，单位为分(必须)
	_totalFee int64
	// 酒店外部编码
	_hotelCode string
	// 外部订单号（必选），阿里旅行会根据此值进行幂等性校验
	_outOid string
	// 房间数
	_roomNum int64
}

// New

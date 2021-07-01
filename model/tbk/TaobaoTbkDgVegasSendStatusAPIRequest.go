package tbk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTbkDgVegasSendStatusAPIRequest
淘宝客-推广者-超级红包领取状态查询 API请求
taobao.tbk.dg.vegas.send.status

淘宝客传入用户标识的信息，查询该用户是否有当前阶段待核销的红包（淘客接入前需签署协议 https://pub.alimama.com/fourth/protocol/common.htm?key=hangye_laxin） */
type TaobaoTbkDgVegasSendStatusAPIRequest struct {
	model.Params
	// 渠道管理id
	_relationId string
	// 会员运营id
	_specialId string
	// 加密后的值(ALIPAY_ID除外)，需用MD5加密，32位小写
	_deviceValue string
	// 入参类型(该模式下返回的结果为模糊匹配结果，和实际情况可能存在误差)： 1. IMEI 2. IDFA 3. OAID 4. MOBILE 5. ALIPAY_ID
	_deviceType string
	// thor平台业务码， 1：coupon 超红
	_thorBizCode string
	// 媒体pid
	_pid string
}

// New

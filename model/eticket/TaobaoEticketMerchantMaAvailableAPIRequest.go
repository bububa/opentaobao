package eticket

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoEticketMerchantMaAvailableAPIRequest
电子凭证核销前校验接口 API请求
taobao.eticket.merchant.ma.available

商家验码之前的调用接口，用来判断是否可以进行核销操作 */
type TaobaoEticketMerchantMaAvailableAPIRequest struct {
	model.Params
	// 业务类型
	_bizType int64
	// 需要被核销的码
	_code string
	// 核销份数
	_consumeNum int64
	// 业务id（订单号）
	_outerId string
	// 机具编号
	_posId string
	// 核销序列号，需要保证唯一
	_serialNum string
	// 需要跟发码通知获取到的参数一致
	_token string
}

// New

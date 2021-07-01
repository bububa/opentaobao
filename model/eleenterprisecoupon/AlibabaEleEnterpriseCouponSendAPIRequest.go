package eleenterprisecoupon

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaEleEnterpriseCouponSendAPIRequest
发放优惠券 API请求
alibaba.ele.enterprise.coupon.send

发放优惠券 */
type AlibabaEleEnterpriseCouponSendAPIRequest struct {
	model.Params
	// 手机号
	_phone string
	// 固定值，申请获取
	_channel string
	// 纬度
	_latitude string
	// 经度
	_longitude string
	// 客户端IP地址
	_ip string
	// 客户端User-Agent信息
	_userAgent string
	// 批次,同一个批次号只会发券一次，后续用同一个批次号的请求会返回上次发的券(幂等)
	_batchNo string
	// 设备ID
	_deviceId string
}

// New

package cainiaohandover

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* CainiaoGlobalHandoverCloudprintGetAPIRequest
获取面单云打印数据 API请求
cainiao.global.handover.cloudprint.get

提供给ISV通过该接口获取面单云打印数据 */
type CainiaoGlobalHandoverCloudprintGetAPIRequest struct {
	model.Params
	// 用户信息
	_userInfo *UserInfoDto
	// 大包运单号
	_trackingNumber string
	// ISV名称，ISV：ISV-ISV英文或拼音名称、商家ERP：SELLER-商家英文或拼音名称
	_client string
	// 大包物流单LP号
	_orderCode string
	// 多语言
	_locale string
}

// New

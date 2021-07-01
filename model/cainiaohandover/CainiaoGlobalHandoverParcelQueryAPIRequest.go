package cainiaohandover

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* CainiaoGlobalHandoverParcelQueryAPIRequest
获取交接单小包信息 API请求
cainiao.global.handover.parcel.query

提供给ISV通过该接口查询小包信息 */
type CainiaoGlobalHandoverParcelQueryAPIRequest struct {
	model.Params
	// 用户信息
	_userInfo *UserInfoDto
	// 多语言
	_locale string
	// 小包的物流订单号,和小包的国际运单号参数任选其一即可
	_orderCode string
	// 小包的国际运单号，和小包的物流订单号参数任选其一即可
	_trackingNumber string
	// 客户端名称，ISV：ISV-ISV英文或拼音名称、商家ERP：SELLER-商家英文或拼音名称
	_client string
}

// New

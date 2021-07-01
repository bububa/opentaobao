package cainiaohandover

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* CainiaoGlobalHandoverContentQueryAPIRequest
查询大包详情 API请求
cainiao.global.handover.content.query

查询大包详情 */
type CainiaoGlobalHandoverContentQueryAPIRequest struct {
	model.Params
	// 用户信息
	_userInfo *UserInfoDto
	// 交接物运单号，和交接物物流订单编码参数任选其一即可
	_trackingNumber string
	// 客户端名称，ISV：ISV-ISV英文或拼音名称、商家ERP：SELLER-商家英文或拼音名称
	_client string
	// 交接物物流订单编码,和交接物运单号参数可以任选其一即可
	_orderCode string
	// 多语言
	_locale string
}

// New

package cainiaohandover

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* CainiaoGlobalHandoverCancelAPIRequest
取消交接单 API请求
cainiao.global.handover.cancel

提供给ISV通过该接口取消交接单 */
type CainiaoGlobalHandoverCancelAPIRequest struct {
	model.Params
	// 系统自动生成
	_userInfo *UserInfoDto
	// 要取消的交接物运单号，即大包运单号
	_trackingNumber string
	// 要取消的交接单id
	_handoverOrderId int64
	// 要取消的交接物id，即大包id
	_handoverContentId int64
	// ISV名称，ISV：ISV-ISV英文或拼音名称、商家ERP：SELLER-商家英文或拼音名称
	_client string
	// 多语言
	_locale string
}

// New

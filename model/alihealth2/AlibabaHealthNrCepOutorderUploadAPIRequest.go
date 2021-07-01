package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaHealthNrCepOutorderUploadAPIRequest
线上订单收货验收单、出入库单据生成接口 API请求
alibaba.health.nr.cep.outorder.upload

线上订单收货验收单、出入库单据生成接口 */
type AlibabaHealthNrCepOutorderUploadAPIRequest struct {
	model.Params
	// 出库单对象
	_topWarOutDto *TopWarOutDto
}

// New

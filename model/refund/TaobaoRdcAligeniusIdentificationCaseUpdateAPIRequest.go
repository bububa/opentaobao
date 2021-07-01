package refund

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoRdcAligeniusIdentificationCaseUpdateAPIRequest
鉴定工单信息同步 API请求
taobao.rdc.aligenius.identification.case.update

同步商家鉴定工单信息 */
type TaobaoRdcAligeniusIdentificationCaseUpdateAPIRequest struct {
	model.Params
	// 请求参数
	_param *SyncIdentifyRefundCaseDto
}

// New

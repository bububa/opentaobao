package refund

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoRdcAligeniusIdentificationCaseResultUpdateAPIRequest
鉴定工单结果同步 API请求
taobao.rdc.aligenius.identification.case.result.update

同步鉴定工单结果信息 */
type TaobaoRdcAligeniusIdentificationCaseResultUpdateAPIRequest struct {
	model.Params
	// 请求参数
	_param *SyncIdentifyRefundCaseResultDto
}

// New

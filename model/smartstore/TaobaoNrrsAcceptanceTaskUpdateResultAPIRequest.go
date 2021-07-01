package smartstore

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoNrrsAcceptanceTaskUpdateResultAPIRequest
更新验收任务支付宝结果 API请求
taobao.nrrs.acceptance.task.updateResult

智慧门店商家验收任务检查相关接口-更新支付宝的验收结果。 */
type TaobaoNrrsAcceptanceTaskUpdateResultAPIRequest struct {
	model.Params
	// 任务ID
	_taskId string
	// 系统自动生成
	_alipayResultList []AlipayCheckResult
}

// New

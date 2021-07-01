package tmallcar

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallAliautoAutofinanceLoanReceiveAPIRequest
接收放款结果通知 API请求
tmall.aliauto.autofinance.loan.receive

天猫汽车的金融业务场景中，需要接收外部ISV对用户支用放款的通知结果 */
type TmallAliautoAutofinanceLoanReceiveAPIRequest struct {
	model.Params
	// 接收外部金融结构的放款结果通知参数
	_loanReceiveDto *LoanReceiveDto
}

// New

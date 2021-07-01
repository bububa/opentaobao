package servicecenter

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoFuwuSpConfirmApplyAPIRequest
内购服务确认单申请接口 API请求
taobao.fuwu.sp.confirm.apply

isv能通过该接口发起确认申请单 */
type TaobaoFuwuSpConfirmApplyAPIRequest struct {
	model.Params
	// 确认单申请
	_paramIncomeConfirmDTO *IncomeConfirmDto
}

// New

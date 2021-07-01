package tmallsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallServicecenterWorkcardRefuseAPIRequest
买家拒收 API请求
tmall.servicecenter.workcard.refuse

买家拒收通知接口 */
type TmallServicecenterWorkcardRefuseAPIRequest struct {
	model.Params
	// 买家拒收信息
	_buyerRefuseAcceptRequest *BuyerRefuseAcceptRequest
}

// New

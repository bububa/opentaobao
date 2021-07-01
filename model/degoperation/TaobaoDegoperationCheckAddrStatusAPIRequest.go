package degoperation

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoDegoperationCheckAddrStatusAPIRequest
地址 API请求
taobao.degoperation.check.addr.status

激励 */
type TaobaoDegoperationCheckAddrStatusAPIRequest struct {
	model.Params
	// 奖品唯一标识
	_sequenceNo int64
}

// New

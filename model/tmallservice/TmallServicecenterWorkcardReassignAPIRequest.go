package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallServicecenterWorkcardReassignAPIRequest
工单改派门店 API请求
tmall.servicecenter.workcard.reassign

工单改派门店 */
type TmallServicecenterWorkcardReassignAPIRequest struct {
	model.Params
	// 请求入参
	_reassignStoreRequest *ReassignStoreRequest
}

// New

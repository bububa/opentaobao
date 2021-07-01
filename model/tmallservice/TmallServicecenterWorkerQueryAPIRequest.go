package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallServicecenterWorkerQueryAPIRequest
工人信息查询 API请求
tmall.servicecenter.worker.query

查询服务商对应的工人信息 */
type TmallServicecenterWorkerQueryAPIRequest struct {
	model.Params
	// 身份证号
	_identityId string
}

// New

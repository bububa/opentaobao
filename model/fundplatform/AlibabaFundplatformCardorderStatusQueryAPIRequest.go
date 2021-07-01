package fundplatform

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaFundplatformCardorderStatusQueryAPIRequest
查询制卡商制卡进度 API请求
alibaba.fundplatform.cardorder.status.query

当通知制卡商进行制卡后，其制卡流程会比较长，若长时间未反馈当前制卡进展，则需要使用该接口来向制卡商发起进度查询。 */
type AlibabaFundplatformCardorderStatusQueryAPIRequest struct {
	model.Params
	// 请求结构体
	_request *AlibabaFundplatformCardorderStatusQueryStruct
}

// New

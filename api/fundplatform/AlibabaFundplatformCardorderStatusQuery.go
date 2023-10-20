package fundplatform

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fundplatform"
)

// AlibabaFundplatformCardorderStatusQuery 查询制卡商制卡进度
// alibaba.fundplatform.cardorder.status.query
//
// 当通知制卡商进行制卡后，其制卡流程会比较长，若长时间未反馈当前制卡进展，则需要使用该接口来向制卡商发起进度查询。
func AlibabaFundplatformCardorderStatusQuery(clt *core.SDKClient, req *fundplatform.AlibabaFundplatformCardorderStatusQueryAPIRequest, resp *fundplatform.AlibabaFundplatformCardorderStatusQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

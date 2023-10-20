package seaking

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/seaking"
)

// AlibabaSeakingTitlerewriteSubmit 提交标题改写任务
// alibaba.seaking.titlerewrite.submit
//
// 提交标题改写任务
func AlibabaSeakingTitlerewriteSubmit(clt *core.SDKClient, req *seaking.AlibabaSeakingTitlerewriteSubmitAPIRequest, resp *seaking.AlibabaSeakingTitlerewriteSubmitAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

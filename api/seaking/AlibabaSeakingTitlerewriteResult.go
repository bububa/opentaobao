package seaking

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/seaking"
)

// AlibabaSeakingTitlerewriteResult 获取标题改写任务结果
// alibaba.seaking.titlerewrite.result
//
// 获取标题改写任务结果
func AlibabaSeakingTitlerewriteResult(clt *core.SDKClient, req *seaking.AlibabaSeakingTitlerewriteResultAPIRequest, resp *seaking.AlibabaSeakingTitlerewriteResultAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

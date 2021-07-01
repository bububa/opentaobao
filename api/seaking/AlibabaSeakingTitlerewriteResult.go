package seaking

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/seaking"
)

/* AlibabaSeakingTitlerewriteResult
获取标题改写任务结果
alibaba.seaking.titlerewrite.result

获取标题改写任务结果 */
func AlibabaSeakingTitlerewriteResult(clt *core.SDKClient, req *seaking.AlibabaSeakingTitlerewriteResultAPIRequest, session string) (*seaking.AlibabaSeakingTitlerewriteResultAPIResponse, error) {
	var resp seaking.AlibabaSeakingTitlerewriteResultAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

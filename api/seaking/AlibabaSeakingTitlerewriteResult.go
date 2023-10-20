package seaking

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/seaking"
)

// Alibabaseakingtitlerewriteresult 获取标题改写任务结果
// alibaba.seaking.titlerewrite.result
//
// 获取标题改写任务结果
func Alibabaseakingtitlerewriteresult(clt *core.SDKClient, req *seaking.AlibabaseakingtitlerewriteresultAPIRequest, session string) (*seaking.AlibabaseakingtitlerewriteresultAPIResponse, error) {
	var resp seaking.AlibabaseakingtitlerewriteresultAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

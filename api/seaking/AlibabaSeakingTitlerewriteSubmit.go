package seaking

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/seaking"
)

// Alibabaseakingtitlerewritesubmit 提交标题改写任务
// alibaba.seaking.titlerewrite.submit
//
// 提交标题改写任务
func Alibabaseakingtitlerewritesubmit(clt *core.SDKClient, req *seaking.AlibabaseakingtitlerewritesubmitAPIRequest, session string) (*seaking.AlibabaseakingtitlerewritesubmitAPIResponse, error) {
	var resp seaking.AlibabaseakingtitlerewritesubmitAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

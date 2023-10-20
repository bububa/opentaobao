package seaking

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/seaking"
)

// Alibabaseakingtaskreport 跳转任务发布成功商品ID回传
// alibaba.seaking.task.report
//
// 跳转任务发布成功商品ID回传
func Alibabaseakingtaskreport(clt *core.SDKClient, req *seaking.AlibabaseakingtaskreportAPIRequest, session string) (*seaking.AlibabaseakingtaskreportAPIResponse, error) {
	var resp seaking.AlibabaseakingtaskreportAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

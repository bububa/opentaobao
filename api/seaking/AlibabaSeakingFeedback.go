package seaking

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/seaking"
)

// AlibabaSeakingFeedback API服务发布成功商品ID回传
// alibaba.seaking.feedback
//
// API服务发布成功商品ID回传，用于跟进商品id后续的使用情况
func AlibabaSeakingFeedback(clt *core.SDKClient, req *seaking.AlibabaSeakingFeedbackAPIRequest, resp *seaking.AlibabaSeakingFeedbackAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

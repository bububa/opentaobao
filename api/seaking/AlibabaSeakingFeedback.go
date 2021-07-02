package seaking

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/seaking"
)

// AlibabaSeakingFeedback API服务发布成功商品ID回传
// alibaba.seaking.feedback
//
// API服务发布成功商品ID回传，用于跟进商品id后续的使用情况
func AlibabaSeakingFeedback(clt *core.SDKClient, req *seaking.AlibabaSeakingFeedbackAPIRequest, session string) (*seaking.AlibabaSeakingFeedbackAPIResponse, error) {
	var resp seaking.AlibabaSeakingFeedbackAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

package seaking

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/seaking"
)

// Alibabaseakingfeedback API服务发布成功商品ID回传
// alibaba.seaking.feedback
//
// API服务发布成功商品ID回传，用于跟进商品id后续的使用情况
func Alibabaseakingfeedback(clt *core.SDKClient, req *seaking.AlibabaseakingfeedbackAPIRequest, session string) (*seaking.AlibabaseakingfeedbackAPIResponse, error) {
	var resp seaking.AlibabaseakingfeedbackAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

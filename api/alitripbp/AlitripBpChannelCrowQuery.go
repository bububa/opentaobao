package alitripbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripbp"
)

// AlitripBpChannelCrowQuery 人群匹配
// alitrip.bp.channel.crow.query
//
// 判断用户是否在圈选的人群中
func AlitripBpChannelCrowQuery(clt *core.SDKClient, req *alitripbp.AlitripBpChannelCrowQueryAPIRequest, session string) (*alitripbp.AlitripBpChannelCrowQueryAPIResponse, error) {
	var resp alitripbp.AlitripBpChannelCrowQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

package alitripbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripbp"
)

// Alitripbpchannelcrowquery 人群匹配
// alitrip.bp.channel.crow.query
//
// 判断用户是否在圈选的人群中
func Alitripbpchannelcrowquery(clt *core.SDKClient, req *alitripbp.AlitripbpchannelcrowqueryAPIRequest, session string) (*alitripbp.AlitripbpchannelcrowqueryAPIResponse, error) {
	var resp alitripbp.AlitripbpchannelcrowqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

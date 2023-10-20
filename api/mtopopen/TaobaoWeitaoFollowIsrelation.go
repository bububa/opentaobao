package mtopopen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mtopopen"
)

// TaobaoWeitaoFollowIsrelation 微淘是否关注
// taobao.weitao.follow.isrelation
//
// 判断用户是否关注对应的公共账号
func TaobaoWeitaoFollowIsrelation(clt *core.SDKClient, req *mtopopen.TaobaoWeitaoFollowIsrelationAPIRequest, session string) (*mtopopen.TaobaoWeitaoFollowIsrelationAPIResponse, error) {
	var resp mtopopen.TaobaoWeitaoFollowIsrelationAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

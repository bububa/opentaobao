package mtopopen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mtopopen"
)

// Taobaoweitaofollowisrelation 微淘是否关注
// taobao.weitao.follow.isrelation
//
// 判断用户是否关注对应的公共账号
func Taobaoweitaofollowisrelation(clt *core.SDKClient, req *mtopopen.TaobaoweitaofollowisrelationAPIRequest, session string) (*mtopopen.TaobaoweitaofollowisrelationAPIResponse, error) {
	var resp mtopopen.TaobaoweitaofollowisrelationAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

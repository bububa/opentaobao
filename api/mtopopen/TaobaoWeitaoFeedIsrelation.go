package mtopopen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mtopopen"
)

// Taobaoweitaofeedisrelation 是否关注
// taobao.weitao.feed.isrelation
//
// 判断用户是否关注对应的公共账号
func Taobaoweitaofeedisrelation(clt *core.SDKClient, req *mtopopen.TaobaoweitaofeedisrelationAPIRequest, session string) (*mtopopen.TaobaoweitaofeedisrelationAPIResponse, error) {
	var resp mtopopen.TaobaoweitaofeedisrelationAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

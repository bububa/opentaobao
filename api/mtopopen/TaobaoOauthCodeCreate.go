package mtopopen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mtopopen"
)

// Taobaooauthcodecreate 淘宝OauthCode颁发
// taobao.oauth.code.create
//
// 手淘无线开放的oauthCode颁发接口
func Taobaooauthcodecreate(clt *core.SDKClient, req *mtopopen.TaobaooauthcodecreateAPIRequest, session string) (*mtopopen.TaobaooauthcodecreateAPIResponse, error) {
	var resp mtopopen.TaobaooauthcodecreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

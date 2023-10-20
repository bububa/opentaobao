package yunosappstore

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/yunosappstore"
)

// Yunosappstoreopengetads 获取外投广告
// yunos.appstore.open.getads
//
// 将广告外投给外部合作伙伴
func Yunosappstoreopengetads(clt *core.SDKClient, req *yunosappstore.YunosappstoreopengetadsAPIRequest, session string) (*yunosappstore.YunosappstoreopengetadsAPIResponse, error) {
	var resp yunosappstore.YunosappstoreopengetadsAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

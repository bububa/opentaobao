package yunosappstore

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/yunosappstore"
)

// YunosAppstoreOpenGetads 获取外投广告
// yunos.appstore.open.getads
//
// 将广告外投给外部合作伙伴
func YunosAppstoreOpenGetads(clt *core.SDKClient, req *yunosappstore.YunosAppstoreOpenGetadsAPIRequest, session string) (*yunosappstore.YunosAppstoreOpenGetadsAPIResponse, error) {
	var resp yunosappstore.YunosAppstoreOpenGetadsAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

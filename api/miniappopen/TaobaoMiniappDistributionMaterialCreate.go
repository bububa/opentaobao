package miniappopen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/miniappopen"
)

// Taobaominiappdistributionmaterialcreate 小程序投放--新建投放素材
// taobao.miniapp.distribution.material.create
//
// 为可投放的小程序，增加入口的素材信息，比如图片、引导文案等等。
func Taobaominiappdistributionmaterialcreate(clt *core.SDKClient, req *miniappopen.TaobaominiappdistributionmaterialcreateAPIRequest, session string) (*miniappopen.TaobaominiappdistributionmaterialcreateAPIResponse, error) {
	var resp miniappopen.TaobaominiappdistributionmaterialcreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

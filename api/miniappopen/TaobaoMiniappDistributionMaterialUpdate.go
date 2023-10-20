package miniappopen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/miniappopen"
)

// Taobaominiappdistributionmaterialupdate 小程序投放 --- 更新投放素材
// taobao.miniapp.distribution.material.update
//
// 修改已录入的投放素材信息。
func Taobaominiappdistributionmaterialupdate(clt *core.SDKClient, req *miniappopen.TaobaominiappdistributionmaterialupdateAPIRequest, session string) (*miniappopen.TaobaominiappdistributionmaterialupdateAPIResponse, error) {
	var resp miniappopen.TaobaominiappdistributionmaterialupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

package miniappopen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/miniappopen"
)

// Taobaominiappdistributionmaterialdelete 小程序投放 --- 删除投放素材
// taobao.miniapp.distribution.material.delete
//
// 删除已录入的投放入口素材信息。
func Taobaominiappdistributionmaterialdelete(clt *core.SDKClient, req *miniappopen.TaobaominiappdistributionmaterialdeleteAPIRequest, session string) (*miniappopen.TaobaominiappdistributionmaterialdeleteAPIResponse, error) {
	var resp miniappopen.TaobaominiappdistributionmaterialdeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

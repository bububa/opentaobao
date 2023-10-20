package miniappopen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/miniappopen"
)

// Taobaominiappdistributionmaterialget 小程序投放---读取投放入口素材信息
// taobao.miniapp.distribution.material.get
//
// 读取已录入的投放入口素材信息。
func Taobaominiappdistributionmaterialget(clt *core.SDKClient, req *miniappopen.TaobaominiappdistributionmaterialgetAPIRequest, session string) (*miniappopen.TaobaominiappdistributionmaterialgetAPIResponse, error) {
	var resp miniappopen.TaobaominiappdistributionmaterialgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

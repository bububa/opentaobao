package usergrowth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/usergrowth2"
)

// TaobaoUsergrowthAdMaterialQuery 素材审核结果查询
// taobao.usergrowth.ad.material.query
//
// 素材审核结果查询
func TaobaoUsergrowthAdMaterialQuery(clt *core.SDKClient, req *usergrowth2.TaobaoUsergrowthAdMaterialQueryAPIRequest, resp *usergrowth2.TaobaoUsergrowthAdMaterialQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

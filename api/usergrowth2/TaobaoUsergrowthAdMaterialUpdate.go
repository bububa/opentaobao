package usergrowth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/usergrowth2"
)

// TaobaoUsergrowthAdMaterialUpdate 素材更新
// taobao.usergrowth.ad.material.update
//
// 素材更新
func TaobaoUsergrowthAdMaterialUpdate(clt *core.SDKClient, req *usergrowth2.TaobaoUsergrowthAdMaterialUpdateAPIRequest, resp *usergrowth2.TaobaoUsergrowthAdMaterialUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

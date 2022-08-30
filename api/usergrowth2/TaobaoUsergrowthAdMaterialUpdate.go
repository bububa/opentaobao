package usergrowth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/usergrowth2"
)

// TaobaoUsergrowthAdMaterialUpdate 素材更新
// taobao.usergrowth.ad.material.update
//
// 素材更新
func TaobaoUsergrowthAdMaterialUpdate(clt *core.SDKClient, req *usergrowth2.TaobaoUsergrowthAdMaterialUpdateAPIRequest, session string) (*usergrowth2.TaobaoUsergrowthAdMaterialUpdateAPIResponse, error) {
	var resp usergrowth2.TaobaoUsergrowthAdMaterialUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

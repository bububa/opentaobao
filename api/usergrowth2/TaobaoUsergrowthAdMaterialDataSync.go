package usergrowth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/usergrowth2"
)

// TaobaoUsergrowthAdMaterialDataSync 素材投放效果数据回传
// taobao.usergrowth.ad.material.data.sync
//
// 创意维度广告效果数据回传
func TaobaoUsergrowthAdMaterialDataSync(clt *core.SDKClient, req *usergrowth2.TaobaoUsergrowthAdMaterialDataSyncAPIRequest, session string) (*usergrowth2.TaobaoUsergrowthAdMaterialDataSyncAPIResponse, error) {
	var resp usergrowth2.TaobaoUsergrowthAdMaterialDataSyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

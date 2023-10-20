package usergrowth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/usergrowth2"
)

// TaobaoUsergrowthAdMediaDataSync 媒体资源位投放效果数据回传
// taobao.usergrowth.ad.media.data.sync
//
// 创意维度广告效果数据回传
func TaobaoUsergrowthAdMediaDataSync(clt *core.SDKClient, req *usergrowth2.TaobaoUsergrowthAdMediaDataSyncAPIRequest, session string) (*usergrowth2.TaobaoUsergrowthAdMediaDataSyncAPIResponse, error) {
	var resp usergrowth2.TaobaoUsergrowthAdMediaDataSyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

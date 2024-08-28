package iot

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iot"
)

// AlibabaAilabsAligenieTracksSearch 查询音频
// alibaba.ailabs.aligenie.tracks.search
//
// 搜索类目下的音频信息
func AlibabaAilabsAligenieTracksSearch(ctx context.Context, clt *core.SDKClient, req *iot.AlibabaAilabsAligenieTracksSearchAPIRequest, resp *iot.AlibabaAilabsAligenieTracksSearchAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

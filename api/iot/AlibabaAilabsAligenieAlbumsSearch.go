package iot

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iot"
)

// AlibabaAilabsAligenieAlbumsSearch 查询专辑
// alibaba.ailabs.aligenie.albums.search
//
// 搜索类目下的专辑信息
func AlibabaAilabsAligenieAlbumsSearch(ctx context.Context, clt *core.SDKClient, req *iot.AlibabaAilabsAligenieAlbumsSearchAPIRequest, resp *iot.AlibabaAilabsAligenieAlbumsSearchAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

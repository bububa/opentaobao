package mos

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mos"
)

// AlibabaMosStoreRecordscreenpointinfo 云屏埋点数据记录接口
// alibaba.mos.store.recordscreenpointinfo
//
// 记录云屏埋点数据
func AlibabaMosStoreRecordscreenpointinfo(ctx context.Context, clt *core.SDKClient, req *mos.AlibabaMosStoreRecordscreenpointinfoAPIRequest, resp *mos.AlibabaMosStoreRecordscreenpointinfoAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

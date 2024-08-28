package idleparttime

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idleparttime"
)

// AlibabaIdleParttimeSynclog 兼职同步日志
// alibaba.idle.parttime.synclog
//
// 提供给供应商查询的接口
func AlibabaIdleParttimeSynclog(ctx context.Context, clt *core.SDKClient, req *idleparttime.AlibabaIdleParttimeSynclogAPIRequest, resp *idleparttime.AlibabaIdleParttimeSynclogAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

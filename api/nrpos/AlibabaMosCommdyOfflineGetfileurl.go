package nrpos

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nrpos"
)

// AlibabaMosCommdyOfflineGetfileurl 去前置机pos商品离线文件下载地址查询接口
// alibaba.mos.commdy.offline.getfileurl
//
// 去前置机-pos查询离线文件下载地址接口
func AlibabaMosCommdyOfflineGetfileurl(ctx context.Context, clt *core.SDKClient, req *nrpos.AlibabaMosCommdyOfflineGetfileurlAPIRequest, resp *nrpos.AlibabaMosCommdyOfflineGetfileurlAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

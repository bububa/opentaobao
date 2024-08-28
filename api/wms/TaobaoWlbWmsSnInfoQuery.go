package wms

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wms"
)

// TaobaoWlbWmsSnInfoQuery 查询单据序列号信息
// taobao.wlb.wms.sn.info.query
//
// 查询仓库作业的各类单据记录的Sn信息
func TaobaoWlbWmsSnInfoQuery(ctx context.Context, clt *core.SDKClient, req *wms.TaobaoWlbWmsSnInfoQueryAPIRequest, resp *wms.TaobaoWlbWmsSnInfoQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

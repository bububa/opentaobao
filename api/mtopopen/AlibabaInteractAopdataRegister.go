package mtopopen

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mtopopen"
)

// AlibabaInteractAopdataRegister 资源位数据推送接口
// alibaba.interact.aopdata.register
//
// 提供给isv，查询以及推送浮层资源位的三方活动数据
func AlibabaInteractAopdataRegister(ctx context.Context, clt *core.SDKClient, req *mtopopen.AlibabaInteractAopdataRegisterAPIRequest, resp *mtopopen.AlibabaInteractAopdataRegisterAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

package alihouse

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeVrSync VR关系数据同步
// alibaba.alihouse.newhome.vr.sync
//
// 对接易居VR关系数据迁移
func AlibabaAlihouseNewhomeVrSync(ctx context.Context, clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeVrSyncAPIRequest, resp *alihouse.AlibabaAlihouseNewhomeVrSyncAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

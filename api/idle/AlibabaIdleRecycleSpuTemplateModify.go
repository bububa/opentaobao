package idle

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// AlibabaIdleRecycleSpuTemplateModify 闲鱼接收回收商spu模板挂载信息
// alibaba.idle.recycle.spu.template.modify
//
// 闲鱼接收回收商spu模板挂载信息
func AlibabaIdleRecycleSpuTemplateModify(ctx context.Context, clt *core.SDKClient, req *idle.AlibabaIdleRecycleSpuTemplateModifyAPIRequest, resp *idle.AlibabaIdleRecycleSpuTemplateModifyAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

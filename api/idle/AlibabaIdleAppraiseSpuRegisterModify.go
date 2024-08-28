package idle

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// AlibabaIdleAppraiseSpuRegisterModify 验货宝服务商spu挂载
// alibaba.idle.appraise.spu.register.modify
//
// 闲鱼接收回收商spu模板挂载信息
func AlibabaIdleAppraiseSpuRegisterModify(ctx context.Context, clt *core.SDKClient, req *idle.AlibabaIdleAppraiseSpuRegisterModifyAPIRequest, resp *idle.AlibabaIdleAppraiseSpuRegisterModifyAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

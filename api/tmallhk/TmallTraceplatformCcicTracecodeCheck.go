package tmallhk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallhk"
)

// TmallTraceplatformCcicTracecodeCheck ccic校验溯源码
// tmall.traceplatform.ccic.tracecode.check
//
// 天猫国际溯源业务，需要将溯源码校验的功能输出到ccic官方主页中以增强溯源码的可信度，故需要提供api给ccic使用以校验溯源码的正确性。
func TmallTraceplatformCcicTracecodeCheck(ctx context.Context, clt *core.SDKClient, req *tmallhk.TmallTraceplatformCcicTracecodeCheckAPIRequest, resp *tmallhk.TmallTraceplatformCcicTracecodeCheckAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

package tmallgenie

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// AlibabaAilabsTmallgenieThirdUnicomShenyanOper 联通神眼注册操作
// alibaba.ailabs.tmallgenie.third.unicom.shenyan.oper
//
// 联通神眼注册操作
func AlibabaAilabsTmallgenieThirdUnicomShenyanOper(ctx context.Context, clt *core.SDKClient, req *tmallgenie.AlibabaAilabsTmallgenieThirdUnicomShenyanOperAPIRequest, resp *tmallgenie.AlibabaAilabsTmallgenieThirdUnicomShenyanOperAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

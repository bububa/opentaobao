package scbp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpEffectAccountDateGet 获取最近报表生成时间
// alibaba.scbp.effect.account.date.get
//
// 获取最近报表生成时间,格式为yyyy-MM-dd
func AlibabaScbpEffectAccountDateGet(ctx context.Context, clt *core.SDKClient, req *scbp.AlibabaScbpEffectAccountDateGetAPIRequest, resp *scbp.AlibabaScbpEffectAccountDateGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

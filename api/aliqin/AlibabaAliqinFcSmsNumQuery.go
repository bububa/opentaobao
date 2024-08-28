package aliqin

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aliqin"
)

// AlibabaAliqinFcSmsNumQuery 短信发送记录查询
// alibaba.aliqin.fc.sms.num.query
//
// 短信发送记录查询。
func AlibabaAliqinFcSmsNumQuery(ctx context.Context, clt *core.SDKClient, req *aliqin.AlibabaAliqinFcSmsNumQueryAPIRequest, resp *aliqin.AlibabaAliqinFcSmsNumQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

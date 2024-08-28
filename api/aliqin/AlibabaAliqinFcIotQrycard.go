package aliqin

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aliqin"
)

// AlibabaAliqinFcIotQrycard 查询终端信息
// alibaba.aliqin.fc.iot.qrycard
//
// 查询终端信息
func AlibabaAliqinFcIotQrycard(ctx context.Context, clt *core.SDKClient, req *aliqin.AlibabaAliqinFcIotQrycardAPIRequest, resp *aliqin.AlibabaAliqinFcIotQrycardAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

package aliqin

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aliqin"
)

// AlibabaAliqinFcIotCardInfo 物联卡信息查询
// alibaba.aliqin.fc.iot.cardInfo
//
// 物联卡信息查询
func AlibabaAliqinFcIotCardInfo(ctx context.Context, clt *core.SDKClient, req *aliqin.AlibabaAliqinFcIotCardInfoAPIRequest, resp *aliqin.AlibabaAliqinFcIotCardInfoAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

package shenjing

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/shenjing"
)

// AlibabaIbShenjingVisitorPadGetqrcodelink pad获取二维码
// alibaba.ib.shenjing.visitor.pad.getqrcodelink
//
// pad获取二维码链接。扫码录入人脸。
func AlibabaIbShenjingVisitorPadGetqrcodelink(ctx context.Context, clt *core.SDKClient, req *shenjing.AlibabaIbShenjingVisitorPadGetqrcodelinkAPIRequest, resp *shenjing.AlibabaIbShenjingVisitorPadGetqrcodelinkAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

package idle

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// AlibabaIdleRentOrderCheckstatusUpload 上传验收结果
// alibaba.idle.rent.order.checkstatus.upload
//
// 上传验收结果
func AlibabaIdleRentOrderCheckstatusUpload(ctx context.Context, clt *core.SDKClient, req *idle.AlibabaIdleRentOrderCheckstatusUploadAPIRequest, resp *idle.AlibabaIdleRentOrderCheckstatusUploadAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

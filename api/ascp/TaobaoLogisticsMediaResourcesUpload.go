package ascp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// TaobaoLogisticsMediaResourcesUpload 图片与视频上传
// taobao.logistics.media.resources.upload
//
// 图片与视频上传
func TaobaoLogisticsMediaResourcesUpload(ctx context.Context, clt *core.SDKClient, req *ascp.TaobaoLogisticsMediaResourcesUploadAPIRequest, resp *ascp.TaobaoLogisticsMediaResourcesUploadAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

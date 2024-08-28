package baichuan

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/baichuan"
)

// AlibabaBaichuanCtgVideoUpload 提供优酷的短视频入淘API
// alibaba.baichuan.ctg.video.upload
//
// 提供优酷的短视频入淘API
func AlibabaBaichuanCtgVideoUpload(ctx context.Context, clt *core.SDKClient, req *baichuan.AlibabaBaichuanCtgVideoUploadAPIRequest, resp *baichuan.AlibabaBaichuanCtgVideoUploadAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

package baichuan

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/baichuan"
)

// AlibabaBaichuanCtgVideoUpload 提供优酷的短视频入淘API
// alibaba.baichuan.ctg.video.upload
//
// 提供优酷的短视频入淘API
func AlibabaBaichuanCtgVideoUpload(clt *core.SDKClient, req *baichuan.AlibabaBaichuanCtgVideoUploadAPIRequest, resp *baichuan.AlibabaBaichuanCtgVideoUploadAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

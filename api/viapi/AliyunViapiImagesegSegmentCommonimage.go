package viapi

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/viapi"
)

// AliyunViapiImagesegSegmentCommonimage 通用分割
// aliyun.viapi.imageseg.segment.commonimage
//
// 识别输入图像中的视觉中心物体轮廓，与背景进行分离，返回分割后的前景物体图（4通道png透明图）。(参数图片/链接必须通过以下方式获取: https://help.aliyun.com/document_detail/155645.html )
func AliyunViapiImagesegSegmentCommonimage(ctx context.Context, clt *core.SDKClient, req *viapi.AliyunViapiImagesegSegmentCommonimageAPIRequest, resp *viapi.AliyunViapiImagesegSegmentCommonimageAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

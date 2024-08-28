package viapi

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/viapi"
)

// AliyunViapiImagesegSegmenthead 头像分割
// aliyun.viapi.imageseg.segmenthead
//
// 输入一张图片，对图中人头区域进行抠图解析，输出人头png透明图。(参数图片/链接必须通过以下方式获取: https://help.aliyun.com/document_detail/155645.html )
func AliyunViapiImagesegSegmenthead(ctx context.Context, clt *core.SDKClient, req *viapi.AliyunViapiImagesegSegmentheadAPIRequest, resp *viapi.AliyunViapiImagesegSegmentheadAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

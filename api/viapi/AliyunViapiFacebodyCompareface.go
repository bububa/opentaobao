package viapi

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/viapi"
)

// AliyunViapiFacebodyCompareface 人脸比对1：1
// aliyun.viapi.facebody.compareface
//
// 基于输入的两张图片，人脸比对服务可检测两张图片中的人脸，并挑选两张图片的最大人脸进行比较，判断是否是同一人；人脸比对服务还返回了这两个人脸的矩形框、比对的置信度，以及不同误识率的置信度阈值。(参数图片/链接必须通过以下方式获取: https://help.aliyun.com/document_detail/155645.html )
func AliyunViapiFacebodyCompareface(ctx context.Context, clt *core.SDKClient, req *viapi.AliyunViapiFacebodyComparefaceAPIRequest, resp *viapi.AliyunViapiFacebodyComparefaceAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

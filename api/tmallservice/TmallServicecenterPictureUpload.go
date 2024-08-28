package tmallservice

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallServicecenterPictureUpload 上传图片
// tmall.servicecenter.picture.upload
//
// 给服务商ERP系统使用，用于上传图片保存在天猫，一般用于工单信息回传时候保存服务商的服务证明信息相关的图片。
func TmallServicecenterPictureUpload(ctx context.Context, clt *core.SDKClient, req *tmallservice.TmallServicecenterPictureUploadAPIRequest, resp *tmallservice.TmallServicecenterPictureUploadAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

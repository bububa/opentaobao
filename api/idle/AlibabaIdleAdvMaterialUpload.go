package idle

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// AlibabaIdleAdvMaterialUpload 闲鱼用户增长素材中心素材上传接口
// alibaba.idle.adv.material.upload
//
// 闲鱼用户增长素材中心素材上传接口
func AlibabaIdleAdvMaterialUpload(ctx context.Context, clt *core.SDKClient, req *idle.AlibabaIdleAdvMaterialUploadAPIRequest, resp *idle.AlibabaIdleAdvMaterialUploadAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

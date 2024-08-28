package mtop

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mtop"
)

// TaobaoMtopUploadTokenGet 获取文件上传授权
// taobao.mtop.upload.token.get
//
// 获取mtop文件上传授权
func TaobaoMtopUploadTokenGet(ctx context.Context, clt *core.SDKClient, req *mtop.TaobaoMtopUploadTokenGetAPIRequest, resp *mtop.TaobaoMtopUploadTokenGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

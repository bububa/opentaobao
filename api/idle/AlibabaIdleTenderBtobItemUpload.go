package idle

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// AlibabaIdleTenderBtobItemUpload 暗拍发布/编辑B2B商品
// alibaba.idle.tender.btob.item.upload
//
// 暗拍发布/编辑B2B商品
func AlibabaIdleTenderBtobItemUpload(ctx context.Context, clt *core.SDKClient, req *idle.AlibabaIdleTenderBtobItemUploadAPIRequest, resp *idle.AlibabaIdleTenderBtobItemUploadAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

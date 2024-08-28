package lbs

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/lbs"
)

// TaobaoLbsMessageUpload lbs数据采集
// taobao.lbs.message.upload
//
// lbs数据采集
func TaobaoLbsMessageUpload(ctx context.Context, clt *core.SDKClient, req *lbs.TaobaoLbsMessageUploadAPIRequest, resp *lbs.TaobaoLbsMessageUploadAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

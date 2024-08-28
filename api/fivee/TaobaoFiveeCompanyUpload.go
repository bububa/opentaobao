package fivee

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fivee"
)

// TaobaoFiveeCompanyUpload 上传商信息接口
// taobao.fivee.company.upload
//
// 资质共享平台上传资质证照
func TaobaoFiveeCompanyUpload(ctx context.Context, clt *core.SDKClient, req *fivee.TaobaoFiveeCompanyUploadAPIRequest, resp *fivee.TaobaoFiveeCompanyUploadAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

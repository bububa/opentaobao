package ihome

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ihome"
)

// TaobaoIhomeAdvancepicUpload ihome图片上传
// taobao.ihome.advancepic.upload
//
// ihome 定制业务编辑器投稿素材上传
func TaobaoIhomeAdvancepicUpload(ctx context.Context, clt *core.SDKClient, req *ihome.TaobaoIhomeAdvancepicUploadAPIRequest, resp *ihome.TaobaoIhomeAdvancepicUploadAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

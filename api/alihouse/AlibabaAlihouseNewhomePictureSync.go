package alihouse

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomePictureSync 图片数据同步
// alibaba.alihouse.newhome.picture.sync
//
// 图片数据同步
func AlibabaAlihouseNewhomePictureSync(ctx context.Context, clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomePictureSyncAPIRequest, resp *alihouse.AlibabaAlihouseNewhomePictureSyncAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

package alihouse

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeShopconfigAstorePreview 天猫好房店铺装修-地址预览
// alibaba.alihouse.newhome.shopconfig.astore.preview
//
// 天猫好房店铺装修-Astore上翻
func AlibabaAlihouseNewhomeShopconfigAstorePreview(ctx context.Context, clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeShopconfigAstorePreviewAPIRequest, resp *alihouse.AlibabaAlihouseNewhomeShopconfigAstorePreviewAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

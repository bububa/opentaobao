package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeShopconfigAstorePreview 天猫好房店铺装修-地址预览
// alibaba.alihouse.newhome.shopconfig.astore.preview
//
// 天猫好房店铺装修-Astore上翻
func AlibabaAlihouseNewhomeShopconfigAstorePreview(clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeShopconfigAstorePreviewAPIRequest, resp *alihouse.AlibabaAlihouseNewhomeShopconfigAstorePreviewAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

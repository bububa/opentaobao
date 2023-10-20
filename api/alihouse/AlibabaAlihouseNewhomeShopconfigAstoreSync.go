package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeShopconfigAstoreSync 天猫好房店铺装修-Astore上翻
// alibaba.alihouse.newhome.shopconfig.astore.sync
//
// 天猫好房店铺装修-Astore上翻
func AlibabaAlihouseNewhomeShopconfigAstoreSync(clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeShopconfigAstoreSyncAPIRequest, resp *alihouse.AlibabaAlihouseNewhomeShopconfigAstoreSyncAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

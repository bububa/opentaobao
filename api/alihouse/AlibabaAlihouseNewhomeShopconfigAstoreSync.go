package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeShopconfigAstoreSync 天猫好房店铺装修-Astore上翻
// alibaba.alihouse.newhome.shopconfig.astore.sync
//
// 天猫好房店铺装修-Astore上翻
func AlibabaAlihouseNewhomeShopconfigAstoreSync(clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeShopconfigAstoreSyncAPIRequest, session string) (*alihouse.AlibabaAlihouseNewhomeShopconfigAstoreSyncAPIResponse, error) {
	var resp alihouse.AlibabaAlihouseNewhomeShopconfigAstoreSyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

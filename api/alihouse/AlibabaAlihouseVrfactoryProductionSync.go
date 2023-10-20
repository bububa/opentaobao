package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseVrfactoryProductionSync vr生产数据上翻
// alibaba.alihouse.vrfactory.production.sync
//
// vr生产数据上翻
func AlibabaAlihouseVrfactoryProductionSync(clt *core.SDKClient, req *alihouse.AlibabaAlihouseVrfactoryProductionSyncAPIRequest, session string) (*alihouse.AlibabaAlihouseVrfactoryProductionSyncAPIResponse, error) {
	var resp alihouse.AlibabaAlihouseVrfactoryProductionSyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

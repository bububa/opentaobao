package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

/* AlibabaAlihouseNewhomeLayoutSync
房通户型数据同步
alibaba.alihouse.newhome.layout.sync

房通户型数据同步 */
func AlibabaAlihouseNewhomeLayoutSync(clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeLayoutSyncAPIRequest, session string) (*alihouse.AlibabaAlihouseNewhomeLayoutSyncAPIResponse, error) {
	var resp alihouse.AlibabaAlihouseNewhomeLayoutSyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

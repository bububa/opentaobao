package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeRegionSync 城区数据同步
// alibaba.alihouse.newhome.region.sync
//
// 城区数据同步
func AlibabaAlihouseNewhomeRegionSync(clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeRegionSyncAPIRequest, session string) (*alihouse.AlibabaAlihouseNewhomeRegionSyncAPIResponse, error) {
	var resp alihouse.AlibabaAlihouseNewhomeRegionSyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

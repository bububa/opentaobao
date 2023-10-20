package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeSupportSync 周边配套数据同步
// alibaba.alihouse.newhome.support.sync
//
// 周边配套数据同步
func AlibabaAlihouseNewhomeSupportSync(clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeSupportSyncAPIRequest, session string) (*alihouse.AlibabaAlihouseNewhomeSupportSyncAPIResponse, error) {
	var resp alihouse.AlibabaAlihouseNewhomeSupportSyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseExistinghomeHouseFeaturesSync 房源标准打标数据同步
// alibaba.alihouse.existinghome.house.features.sync
//
// 房源标准打标数据同步
func AlibabaAlihouseExistinghomeHouseFeaturesSync(clt *core.SDKClient, req *alihouse.AlibabaAlihouseExistinghomeHouseFeaturesSyncAPIRequest, resp *alihouse.AlibabaAlihouseExistinghomeHouseFeaturesSyncAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

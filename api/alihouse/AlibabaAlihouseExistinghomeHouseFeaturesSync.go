package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseExistinghomeHouseFeaturesSync 房源标准打标数据同步
// alibaba.alihouse.existinghome.house.features.sync
//
// 房源标准打标数据同步
func AlibabaAlihouseExistinghomeHouseFeaturesSync(clt *core.SDKClient, req *alihouse.AlibabaAlihouseExistinghomeHouseFeaturesSyncAPIRequest, session string) (*alihouse.AlibabaAlihouseExistinghomeHouseFeaturesSyncAPIResponse, error) {
	var resp alihouse.AlibabaAlihouseExistinghomeHouseFeaturesSyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

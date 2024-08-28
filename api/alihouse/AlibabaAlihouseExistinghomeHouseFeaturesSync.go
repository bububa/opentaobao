package alihouse

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseExistinghomeHouseFeaturesSync 房源标准打标数据同步
// alibaba.alihouse.existinghome.house.features.sync
//
// 房源标准打标数据同步
func AlibabaAlihouseExistinghomeHouseFeaturesSync(ctx context.Context, clt *core.SDKClient, req *alihouse.AlibabaAlihouseExistinghomeHouseFeaturesSyncAPIRequest, resp *alihouse.AlibabaAlihouseExistinghomeHouseFeaturesSyncAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

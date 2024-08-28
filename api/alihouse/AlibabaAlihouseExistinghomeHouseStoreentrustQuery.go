package alihouse

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseExistinghomeHouseStoreentrustQuery 门店委托信息查询
// alibaba.alihouse.existinghome.house.storeentrust.query
//
// 门店委托信息查询
func AlibabaAlihouseExistinghomeHouseStoreentrustQuery(ctx context.Context, clt *core.SDKClient, req *alihouse.AlibabaAlihouseExistinghomeHouseStoreentrustQueryAPIRequest, resp *alihouse.AlibabaAlihouseExistinghomeHouseStoreentrustQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

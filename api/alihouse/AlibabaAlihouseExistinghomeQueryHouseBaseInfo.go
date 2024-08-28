package alihouse

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseExistinghomeQueryHouseBaseInfo 查询房源基本信息
// alibaba.alihouse.existinghome.query.house.base.info
//
// 查询房源基本信息
func AlibabaAlihouseExistinghomeQueryHouseBaseInfo(ctx context.Context, clt *core.SDKClient, req *alihouse.AlibabaAlihouseExistinghomeQueryHouseBaseInfoAPIRequest, resp *alihouse.AlibabaAlihouseExistinghomeQueryHouseBaseInfoAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

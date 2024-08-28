package alihouse

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseExistinghomeQueryHouseTagInfo 活动标查询接口
// alibaba.alihouse.existinghome.query.house.tag.info
//
// 活动标查询接口
func AlibabaAlihouseExistinghomeQueryHouseTagInfo(ctx context.Context, clt *core.SDKClient, req *alihouse.AlibabaAlihouseExistinghomeQueryHouseTagInfoAPIRequest, resp *alihouse.AlibabaAlihouseExistinghomeQueryHouseTagInfoAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

package alihouse

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseExistinghomeActivityQuery 五福活动经纪人获取
// alibaba.alihouse.existinghome.activity.query
//
// 五福活动经纪人获取
func AlibabaAlihouseExistinghomeActivityQuery(ctx context.Context, clt *core.SDKClient, req *alihouse.AlibabaAlihouseExistinghomeActivityQueryAPIRequest, resp *alihouse.AlibabaAlihouseExistinghomeActivityQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

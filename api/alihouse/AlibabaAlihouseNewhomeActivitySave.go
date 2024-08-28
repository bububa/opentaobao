package alihouse

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeActivitySave 新增或者更新销售活动
// alibaba.alihouse.newhome.activity.save
//
// 新增或者更新销售活动
func AlibabaAlihouseNewhomeActivitySave(ctx context.Context, clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeActivitySaveAPIRequest, resp *alihouse.AlibabaAlihouseNewhomeActivitySaveAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

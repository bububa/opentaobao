package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeActivitySave 新增或者更新销售活动
// alibaba.alihouse.newhome.activity.save
//
// 新增或者更新销售活动
func AlibabaAlihouseNewhomeActivitySave(clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeActivitySaveAPIRequest, resp *alihouse.AlibabaAlihouseNewhomeActivitySaveAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

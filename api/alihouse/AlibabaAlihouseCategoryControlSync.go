package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseCategoryControlSync 类目权限上翻
// alibaba.alihouse.category.control.sync
//
// 类目权限上翻
func AlibabaAlihouseCategoryControlSync(clt *core.SDKClient, req *alihouse.AlibabaAlihouseCategoryControlSyncAPIRequest, resp *alihouse.AlibabaAlihouseCategoryControlSyncAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

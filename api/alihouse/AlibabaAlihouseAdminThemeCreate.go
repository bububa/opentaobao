package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseAdminThemeCreate 创建云主题
// alibaba.alihouse.admin.theme.create
//
// 创建云主题
func AlibabaAlihouseAdminThemeCreate(clt *core.SDKClient, req *alihouse.AlibabaAlihouseAdminThemeCreateAPIRequest, resp *alihouse.AlibabaAlihouseAdminThemeCreateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

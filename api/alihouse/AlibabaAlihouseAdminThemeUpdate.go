package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseAdminThemeUpdate 云主题更新
// alibaba.alihouse.admin.theme.update
//
// 云主题更新
func AlibabaAlihouseAdminThemeUpdate(clt *core.SDKClient, req *alihouse.AlibabaAlihouseAdminThemeUpdateAPIRequest, resp *alihouse.AlibabaAlihouseAdminThemeUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

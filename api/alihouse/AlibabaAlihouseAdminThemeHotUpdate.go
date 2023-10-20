package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseAdminThemeHotUpdate 云主题热更新数据集
// alibaba.alihouse.admin.theme.hot.update
//
// 云主题更新
func AlibabaAlihouseAdminThemeHotUpdate(clt *core.SDKClient, req *alihouse.AlibabaAlihouseAdminThemeHotUpdateAPIRequest, resp *alihouse.AlibabaAlihouseAdminThemeHotUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

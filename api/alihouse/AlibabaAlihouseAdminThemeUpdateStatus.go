package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseAdminThemeUpdateStatus 云主题上下架+删除
// alibaba.alihouse.admin.theme.update.status
//
// 云主题上下架+删除
func AlibabaAlihouseAdminThemeUpdateStatus(clt *core.SDKClient, req *alihouse.AlibabaAlihouseAdminThemeUpdateStatusAPIRequest, session string) (*alihouse.AlibabaAlihouseAdminThemeUpdateStatusAPIResponse, error) {
	var resp alihouse.AlibabaAlihouseAdminThemeUpdateStatusAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

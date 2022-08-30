package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseAdminThemeUpdate 云主题更新
// alibaba.alihouse.admin.theme.update
//
// 云主题更新
func AlibabaAlihouseAdminThemeUpdate(clt *core.SDKClient, req *alihouse.AlibabaAlihouseAdminThemeUpdateAPIRequest, session string) (*alihouse.AlibabaAlihouseAdminThemeUpdateAPIResponse, error) {
	var resp alihouse.AlibabaAlihouseAdminThemeUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseAdminThemeCreate 创建云主题
// alibaba.alihouse.admin.theme.create
//
// 创建云主题
func AlibabaAlihouseAdminThemeCreate(clt *core.SDKClient, req *alihouse.AlibabaAlihouseAdminThemeCreateAPIRequest, session string) (*alihouse.AlibabaAlihouseAdminThemeCreateAPIResponse, error) {
	var resp alihouse.AlibabaAlihouseAdminThemeCreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

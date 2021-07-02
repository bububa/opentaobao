package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvscreenAdminCommonOperation 一体机桌面通用接口
// yunos.tvscreen.admin.common.operation
//
// 一体机桌面通用接口
func YunosTvscreenAdminCommonOperation(clt *core.SDKClient, req *tvupadmin.YunosTvscreenAdminCommonOperationAPIRequest, session string) (*tvupadmin.YunosTvscreenAdminCommonOperationAPIResponse, error) {
	var resp tvupadmin.YunosTvscreenAdminCommonOperationAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

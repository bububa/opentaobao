package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// Yunospubadmincommonoperation 内部迎客松通用服务
// yunos.pubadmin.common.operation
//
// 内部迎客松通用服务
func Yunospubadmincommonoperation(clt *core.SDKClient, req *tvupadmin.YunospubadmincommonoperationAPIRequest, session string) (*tvupadmin.YunospubadmincommonoperationAPIResponse, error) {
	var resp tvupadmin.YunospubadmincommonoperationAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

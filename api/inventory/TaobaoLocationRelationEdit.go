package inventory

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/inventory"
)

// Taobaolocationrelationedit 地点关联关系增量编辑
// taobao.location.relation.edit
//
// 地点关联关系增量编辑
func Taobaolocationrelationedit(clt *core.SDKClient, req *inventory.TaobaolocationrelationeditAPIRequest, session string) (*inventory.TaobaolocationrelationeditAPIResponse, error) {
	var resp inventory.TaobaolocationrelationeditAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

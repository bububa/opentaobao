package tblogistics

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tblogistics"
)

// Taobaodeliverytemplateupdate 修改运费模板
// taobao.delivery.template.update
//
// 修改运费模板
func Taobaodeliverytemplateupdate(clt *core.SDKClient, req *tblogistics.TaobaodeliverytemplateupdateAPIRequest, session string) (*tblogistics.TaobaodeliverytemplateupdateAPIResponse, error) {
	var resp tblogistics.TaobaodeliverytemplateupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

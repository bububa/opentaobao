package miniapp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/miniapp"
)

// TaobaoSmartappTableGet 智能应用服务登记工作表数据查询
// taobao.smartapp.table.get
//
// 智能应用服务登记工作表数据查询
func TaobaoSmartappTableGet(clt *core.SDKClient, req *miniapp.TaobaoSmartappTableGetAPIRequest, session string) (*miniapp.TaobaoSmartappTableGetAPIResponse, error) {
	var resp miniapp.TaobaoSmartappTableGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

package miniapp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/miniapp"
)

// TaobaoSmartappTableFulldataGet 智能应用工作表地址查询
// taobao.smartapp.table.fulldata.get
//
// 智能应用工作表地址查询
func TaobaoSmartappTableFulldataGet(clt *core.SDKClient, req *miniapp.TaobaoSmartappTableFulldataGetAPIRequest, session string) (*miniapp.TaobaoSmartappTableFulldataGetAPIResponse, error) {
	var resp miniapp.TaobaoSmartappTableFulldataGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

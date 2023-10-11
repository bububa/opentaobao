package miniapp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/miniapp"
)

// TaobaoSmartappTableAdd 智能应用服务登记工作表数据新增
// taobao.smartapp.table.add
//
// 智能应用服务登记工作表数据新增
func TaobaoSmartappTableAdd(clt *core.SDKClient, req *miniapp.TaobaoSmartappTableAddAPIRequest, session string) (*miniapp.TaobaoSmartappTableAddAPIResponse, error) {
	var resp miniapp.TaobaoSmartappTableAddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

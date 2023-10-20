package wms

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wms"
)

// Taobaowlbwmscainiaobillquery 查询单据列表
// taobao.wlb.wms.cainiao.bill.query
//
// 查询单据列表
func Taobaowlbwmscainiaobillquery(clt *core.SDKClient, req *wms.TaobaowlbwmscainiaobillqueryAPIRequest, session string) (*wms.TaobaowlbwmscainiaobillqueryAPIResponse, error) {
	var resp wms.TaobaowlbwmscainiaobillqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

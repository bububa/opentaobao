package wms

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wms"
)

// TaobaoWlbWmsCainiaoBillQuery 查询单据列表
// taobao.wlb.wms.cainiao.bill.query
//
// 查询单据列表
func TaobaoWlbWmsCainiaoBillQuery(clt *core.SDKClient, req *wms.TaobaoWlbWmsCainiaoBillQueryAPIRequest, session string) (*wms.TaobaoWlbWmsCainiaoBillQueryAPIResponse, error) {
	var resp wms.TaobaoWlbWmsCainiaoBillQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

package waybill

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/waybill"
)

// Taobaowlbwaybillisearch 查询面单服务订购及面单使用情况v1.0
// taobao.wlb.waybill.i.search
//
// 获取发货地&amp;CP开通状态&amp;账户的使用情况
func Taobaowlbwaybillisearch(clt *core.SDKClient, req *waybill.TaobaowlbwaybillisearchAPIRequest, session string) (*waybill.TaobaowlbwaybillisearchAPIResponse, error) {
	var resp waybill.TaobaowlbwaybillisearchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

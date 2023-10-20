package waybill

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/waybill"
)

// Taobaowlbwaybilliquerydetail 查面单号状态v1.0
// taobao.wlb.waybill.i.querydetail
//
// 查看面单号的当前状态，如签收、发货、失效等。
func Taobaowlbwaybilliquerydetail(clt *core.SDKClient, req *waybill.TaobaowlbwaybilliquerydetailAPIRequest, session string) (*waybill.TaobaowlbwaybilliquerydetailAPIResponse, error) {
	var resp waybill.TaobaowlbwaybilliquerydetailAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

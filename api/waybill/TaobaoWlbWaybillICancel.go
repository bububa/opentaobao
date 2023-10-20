package waybill

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/waybill"
)

// Taobaowlbwaybillicancel 商家取消获取的电子面单号v1.0
// taobao.wlb.waybill.i.cancel
//
// 面单号有误需要取消的时候，调用该接口取消获取的电子面单。
func Taobaowlbwaybillicancel(clt *core.SDKClient, req *waybill.TaobaowlbwaybillicancelAPIRequest, session string) (*waybill.TaobaowlbwaybillicancelAPIResponse, error) {
	var resp waybill.TaobaowlbwaybillicancelAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

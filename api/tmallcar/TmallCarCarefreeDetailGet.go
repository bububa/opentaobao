package tmallcar

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallcar"
)

// TmallCarCarefreeDetailGet 查询业务单信息
// tmall.car.carefree.detail.get
//
// 查询业务单信息
func TmallCarCarefreeDetailGet(clt *core.SDKClient, req *tmallcar.TmallCarCarefreeDetailGetAPIRequest, resp *tmallcar.TmallCarCarefreeDetailGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

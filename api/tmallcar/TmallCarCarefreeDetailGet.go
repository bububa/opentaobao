package tmallcar

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallcar"
)

// TmallCarCarefreeDetailGet 查询业务单信息
// tmall.car.carefree.detail.get
//
// 查询业务单信息
func TmallCarCarefreeDetailGet(clt *core.SDKClient, req *tmallcar.TmallCarCarefreeDetailGetAPIRequest, session string) (*tmallcar.TmallCarCarefreeDetailGetAPIResponse, error) {
	var resp tmallcar.TmallCarCarefreeDetailGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

package tmallcar

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallcar"
)

// Tmallcarcarefreedetailget 查询业务单信息
// tmall.car.carefree.detail.get
//
// 查询业务单信息
func Tmallcarcarefreedetailget(clt *core.SDKClient, req *tmallcar.TmallcarcarefreedetailgetAPIRequest, session string) (*tmallcar.TmallcarcarefreedetailgetAPIResponse, error) {
	var resp tmallcar.TmallcarcarefreedetailgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

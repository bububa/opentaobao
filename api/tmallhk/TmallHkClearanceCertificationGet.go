package tmallhk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallhk"
)

// Tmallhkclearancecertificationget 获取订单清关材料实名信息
// tmall.hk.clearance.certification.get
//
// 获取订单清关材料实名信息
func Tmallhkclearancecertificationget(clt *core.SDKClient, req *tmallhk.TmallhkclearancecertificationgetAPIRequest, session string) (*tmallhk.TmallhkclearancecertificationgetAPIResponse, error) {
	var resp tmallhk.TmallhkclearancecertificationgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

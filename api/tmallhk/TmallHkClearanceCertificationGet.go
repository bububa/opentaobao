package tmallhk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallhk"
)

// TmallHkClearanceCertificationGet 获取订单清关材料实名信息
// tmall.hk.clearance.certification.get
//
// 获取订单清关材料实名信息
func TmallHkClearanceCertificationGet(clt *core.SDKClient, req *tmallhk.TmallHkClearanceCertificationGetAPIRequest, session string) (*tmallhk.TmallHkClearanceCertificationGetAPIResponse, error) {
	var resp tmallhk.TmallHkClearanceCertificationGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

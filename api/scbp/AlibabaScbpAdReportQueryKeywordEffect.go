package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// Alibabascbpadreportquerykeywordeffect 关键词报告
// alibaba.scbp.ad.report.query.keyword.effect
//
// 关键词报告
func Alibabascbpadreportquerykeywordeffect(clt *core.SDKClient, req *scbp.AlibabascbpadreportquerykeywordeffectAPIRequest, session string) (*scbp.AlibabascbpadreportquerykeywordeffectAPIResponse, error) {
	var resp scbp.AlibabascbpadreportquerykeywordeffectAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

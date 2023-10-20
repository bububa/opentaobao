package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// Alibabascbpadreportquerysinglekeywordeffect 单个关键词报告
// alibaba.scbp.ad.report.query.single.keyword.effect
//
// 单个关键词报告
func Alibabascbpadreportquerysinglekeywordeffect(clt *core.SDKClient, req *scbp.AlibabascbpadreportquerysinglekeywordeffectAPIRequest, session string) (*scbp.AlibabascbpadreportquerysinglekeywordeffectAPIResponse, error) {
	var resp scbp.AlibabascbpadreportquerysinglekeywordeffectAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

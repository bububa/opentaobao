package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// Alibabascbpadreportgetlasteffectdate 获取最近报表生成时间
// alibaba.scbp.ad.report.get.last.effect.date
//
// 获取最近报表生成时间
func Alibabascbpadreportgetlasteffectdate(clt *core.SDKClient, req *scbp.AlibabascbpadreportgetlasteffectdateAPIRequest, session string) (*scbp.AlibabascbpadreportgetlasteffectdateAPIResponse, error) {
	var resp scbp.AlibabascbpadreportgetlasteffectdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

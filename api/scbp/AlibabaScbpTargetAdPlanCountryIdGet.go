package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// Alibabascbptargetadplancountryidget 定向推广-国家标签ID获取
// alibaba.scbp.target.ad.plan.country.id.get
//
// 定向推广-国家标签ID获取
func Alibabascbptargetadplancountryidget(clt *core.SDKClient, req *scbp.AlibabascbptargetadplancountryidgetAPIRequest, session string) (*scbp.AlibabascbptargetadplancountryidgetAPIResponse, error) {
	var resp scbp.AlibabascbptargetadplancountryidgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

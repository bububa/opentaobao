package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugkytdrgetbyentid 多融通过企业ID得到一个企业的详细信息
// alibaba.alihealth.drug.kyt.dr.getbyentid
//
// 根据企业主键查看企业详细信息
func Alibabaalihealthdrugkytdrgetbyentid(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugkytdrgetbyentidAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugkytdrgetbyentidAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugkytdrgetbyentidAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

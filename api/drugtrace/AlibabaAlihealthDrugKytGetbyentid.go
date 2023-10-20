package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugkytgetbyentid 根据企业主键查看企业详细信息
// alibaba.alihealth.drug.kyt.getbyentid
//
// 根据企业主键查看企业详细信息
func Alibabaalihealthdrugkytgetbyentid(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugkytgetbyentidAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugkytgetbyentidAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugkytgetbyentidAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugkytwesgetbyentid 根据企业主键查看企业详细信息
// alibaba.alihealth.drug.kyt.wes.getbyentid
//
// 根据企业主键查看企业详细信息
func Alibabaalihealthdrugkytwesgetbyentid(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugkytwesgetbyentidAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugkytwesgetbyentidAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugkytwesgetbyentidAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

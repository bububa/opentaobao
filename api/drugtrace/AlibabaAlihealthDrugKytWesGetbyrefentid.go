package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugkytwesgetbyrefentid 根据企业唯一标识查看企业详细信息
// alibaba.alihealth.drug.kyt.wes.getbyrefentid
//
// 根据企业唯一标识查看企业详细信息
func Alibabaalihealthdrugkytwesgetbyrefentid(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugkytwesgetbyrefentidAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugkytwesgetbyrefentidAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugkytwesgetbyrefentidAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

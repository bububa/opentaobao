package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugkytgetbyrefentid 根据企业唯一标识查看企业详细信息
// alibaba.alihealth.drug.kyt.getbyrefentid
//
// 根据企业唯一标识查看企业详细信息
func Alibabaalihealthdrugkytgetbyrefentid(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugkytgetbyrefentidAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugkytgetbyrefentidAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugkytgetbyrefentidAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

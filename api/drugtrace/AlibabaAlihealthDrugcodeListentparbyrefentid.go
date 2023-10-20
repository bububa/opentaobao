package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugcodelistentparbyrefentid 根据企业id获取往来单位
// alibaba.alihealth.drugcode.listentparbyrefentid
//
// 根据企业id获取往来单位
func Alibabaalihealthdrugcodelistentparbyrefentid(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugcodelistentparbyrefentidAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugcodelistentparbyrefentidAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugcodelistentparbyrefentidAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

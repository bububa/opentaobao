package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugcodeListentparbyrefentid 根据企业id获取往来单位
// alibaba.alihealth.drugcode.listentparbyrefentid
//
// 根据企业id获取往来单位
func AlibabaAlihealthDrugcodeListentparbyrefentid(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugcodeListentparbyrefentidAPIRequest, session string) (*drugtrace.AlibabaAlihealthDrugcodeListentparbyrefentidAPIResponse, error) {
	var resp drugtrace.AlibabaAlihealthDrugcodeListentparbyrefentidAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

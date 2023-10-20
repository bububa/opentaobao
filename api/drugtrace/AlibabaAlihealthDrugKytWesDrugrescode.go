package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugkytwesdrugrescode 查询药品码段信息
// alibaba.alihealth.drug.kyt.wes.drugrescode
//
// 查询药品码段信息
func Alibabaalihealthdrugkytwesdrugrescode(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugkytwesdrugrescodeAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugkytwesdrugrescodeAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugkytwesdrugrescodeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

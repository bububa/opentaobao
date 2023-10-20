package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugkytdrugrescode 查询药品码段信息
// alibaba.alihealth.drug.kyt.drugrescode
//
// 查询药品码段信息
func Alibabaalihealthdrugkytdrugrescode(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugkytdrugrescodeAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugkytdrugrescodeAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugkytdrugrescodeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

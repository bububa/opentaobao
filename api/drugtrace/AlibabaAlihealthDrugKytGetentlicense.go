package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugkytgetentlicense 获取企业资质
// alibaba.alihealth.drug.kyt.getentlicense
//
// 获取企业的资质信息。
func Alibabaalihealthdrugkytgetentlicense(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugkytgetentlicenseAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugkytgetentlicenseAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugkytgetentlicenseAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

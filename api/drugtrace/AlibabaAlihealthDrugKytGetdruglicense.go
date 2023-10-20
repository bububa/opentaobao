package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugkytgetdruglicense 获取药品资质信息
// alibaba.alihealth.drug.kyt.getdruglicense
//
// 获取药品的资质信息。
func Alibabaalihealthdrugkytgetdruglicense(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugkytgetdruglicenseAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugkytgetdruglicenseAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugkytgetdruglicenseAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

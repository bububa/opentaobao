package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugkytsaveent 新增往来单位企业
// alibaba.alihealth.drug.kyt.saveent
//
// 新增往来单位企业记录
func Alibabaalihealthdrugkytsaveent(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugkytsaveentAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugkytsaveentAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugkytsaveentAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

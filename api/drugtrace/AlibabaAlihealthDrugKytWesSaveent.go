package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugkytwessaveent 新增往来单位企业记录
// alibaba.alihealth.drug.kyt.wes.saveent
//
// 新增往来单位企业记录
func Alibabaalihealthdrugkytwessaveent(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugkytwessaveentAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugkytwessaveentAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugkytwessaveentAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugkytdrsinglerelation 多融单码关联关系查询
// alibaba.alihealth.drug.kyt.dr.singlerelation
//
// 单码关联关系查询
func Alibabaalihealthdrugkytdrsinglerelation(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugkytdrsinglerelationAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugkytdrsinglerelationAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugkytdrsinglerelationAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

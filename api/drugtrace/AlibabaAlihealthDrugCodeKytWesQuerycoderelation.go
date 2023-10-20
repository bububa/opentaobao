package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugcodekytwesquerycoderelation wes码关联关系查询
// alibaba.alihealth.drug.code.kyt.wes.querycoderelation
//
// 单码关联关系查询，通过一个码查询这个码下的所有子码。（只有做过入库的码，才能能进行查询）
func Alibabaalihealthdrugcodekytwesquerycoderelation(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugcodekytwesquerycoderelationAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugcodekytwesquerycoderelationAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugcodekytwesquerycoderelationAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

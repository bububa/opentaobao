package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugkytdrugdetail 查询药品详细信息
// alibaba.alihealth.drug.kyt.drugdetail
//
// 查询药品详细信息
func Alibabaalihealthdrugkytdrugdetail(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugkytdrugdetailAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugkytdrugdetailAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugkytdrugdetailAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

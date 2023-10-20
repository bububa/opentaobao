package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugkytrelationdetail 关联关系处理详情
// alibaba.alihealth.drug.kyt.relationdetail
//
// 关联关系处理详情
func Alibabaalihealthdrugkytrelationdetail(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugkytrelationdetailAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugkytrelationdetailAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugkytrelationdetailAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

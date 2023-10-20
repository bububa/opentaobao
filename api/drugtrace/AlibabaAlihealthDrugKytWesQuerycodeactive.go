package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugkytwesquerycodeactive 查询码是否激活
// alibaba.alihealth.drug.kyt.wes.querycodeactive
//
// 查询码是否激活
func Alibabaalihealthdrugkytwesquerycodeactive(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugkytwesquerycodeactiveAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugkytwesquerycodeactiveAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugkytwesquerycodeactiveAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

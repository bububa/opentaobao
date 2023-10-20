package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugcodelistcodegov 政府码查询接口
// alibaba.alihealth.drug.code.list.code.gov
//
// 政府码查询接口
func Alibabaalihealthdrugcodelistcodegov(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugcodelistcodegovAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugcodelistcodegovAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugcodelistcodegovAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

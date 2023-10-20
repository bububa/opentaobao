package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugkytgetcodebaseinfo 码的药品信息查询
// alibaba.alihealth.drug.kyt.getcodebaseinfo
//
// 提供根据码查询码基本信息接口
func Alibabaalihealthdrugkytgetcodebaseinfo(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugkytgetcodebaseinfoAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugkytgetcodebaseinfoAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugkytgetcodebaseinfoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

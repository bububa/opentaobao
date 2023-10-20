package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugcodenodenameget 根据码获取机构名称
// alibaba.alihealth.drugcode.nodename.get
//
// 根据码获取机构名称
func Alibabaalihealthdrugcodenodenameget(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugcodenodenamegetAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugcodenodenamegetAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugcodenodenamegetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

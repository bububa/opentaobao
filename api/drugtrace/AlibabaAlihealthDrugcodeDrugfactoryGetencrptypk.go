package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugcodedrugfactorygetencrptypk 获取加密公钥
// alibaba.alihealth.drugcode.drugfactory.getencrptypk
//
// 获取服务端给药厂用来加密的公钥
func Alibabaalihealthdrugcodedrugfactorygetencrptypk(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugcodedrugfactorygetencrptypkAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugcodedrugfactorygetencrptypkAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugcodedrugfactorygetencrptypkAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

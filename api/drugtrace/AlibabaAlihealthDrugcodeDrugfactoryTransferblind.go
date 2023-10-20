package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugcodedrugfactorytransferblind 传输盲底文件
// alibaba.alihealth.drugcode.drugfactory.transferblind
//
// 临床用药试验-传输盲底文件
func Alibabaalihealthdrugcodedrugfactorytransferblind(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugcodedrugfactorytransferblindAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugcodedrugfactorytransferblindAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugcodedrugfactorytransferblindAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

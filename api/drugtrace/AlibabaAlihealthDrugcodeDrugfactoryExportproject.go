package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugcodedrugfactoryexportproject 导出项目和药品目录
// alibaba.alihealth.drugcode.drugfactory.exportproject
//
// 导出临床项目及其药品目录
func Alibabaalihealthdrugcodedrugfactoryexportproject(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugcodedrugfactoryexportprojectAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugcodedrugfactoryexportprojectAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugcodedrugfactoryexportprojectAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

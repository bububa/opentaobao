package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugcodedrugfactoryexportattribute 导出所有项目的药物属性和药品信息
// alibaba.alihealth.drugcode.drugfactory.exportattribute
//
// 导出所有项目的药物属性和药品信息
func Alibabaalihealthdrugcodedrugfactoryexportattribute(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugcodedrugfactoryexportattributeAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugcodedrugfactoryexportattributeAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugcodedrugfactoryexportattributeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugcodedrugfactoryexportcategory 导出临床药品目录
// alibaba.alihealth.drugcode.drugfactory.exportcategory
//
// 导出临床药品目录
func Alibabaalihealthdrugcodedrugfactoryexportcategory(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugcodedrugfactoryexportcategoryAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugcodedrugfactoryexportcategoryAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugcodedrugfactoryexportcategoryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

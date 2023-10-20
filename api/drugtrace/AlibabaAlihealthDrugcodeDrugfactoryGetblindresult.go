package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugcodedrugfactorygetblindresult 获取盲底文件处理结果
// alibaba.alihealth.drugcode.drugfactory.getblindresult
//
// 获取盲底文件处理结果
func Alibabaalihealthdrugcodedrugfactorygetblindresult(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugcodedrugfactorygetblindresultAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugcodedrugfactorygetblindresultAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugcodedrugfactorygetblindresultAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

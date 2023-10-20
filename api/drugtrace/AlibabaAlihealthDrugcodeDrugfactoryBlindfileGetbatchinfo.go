package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugcodedrugfactoryblindfilegetbatchinfo 获取盲底文件中的批次信息
// alibaba.alihealth.drugcode.drugfactory.blindfile.getbatchinfo
//
// 获取盲底文件中的批次信息
func Alibabaalihealthdrugcodedrugfactoryblindfilegetbatchinfo(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugcodedrugfactoryblindfilegetbatchinfoAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugcodedrugfactoryblindfilegetbatchinfoAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugcodedrugfactoryblindfilegetbatchinfoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

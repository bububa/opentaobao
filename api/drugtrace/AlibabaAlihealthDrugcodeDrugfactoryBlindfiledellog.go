package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugcodedrugfactoryblindfiledellog 接收盲底文件删除日志
// alibaba.alihealth.drugcode.drugfactory.blindfiledellog
//
// 临床用药试验-接收盲底文件删除日志
func Alibabaalihealthdrugcodedrugfactoryblindfiledellog(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugcodedrugfactoryblindfiledellogAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugcodedrugfactoryblindfiledellogAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugcodedrugfactoryblindfiledellogAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

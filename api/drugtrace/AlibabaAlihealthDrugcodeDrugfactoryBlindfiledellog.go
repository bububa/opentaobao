package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugcodeDrugfactoryBlindfiledellog 接收盲底文件删除日志
// alibaba.alihealth.drugcode.drugfactory.blindfiledellog
//
// 临床用药试验-接收盲底文件删除日志
func AlibabaAlihealthDrugcodeDrugfactoryBlindfiledellog(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugcodeDrugfactoryBlindfiledellogAPIRequest, resp *drugtrace.AlibabaAlihealthDrugcodeDrugfactoryBlindfiledellogAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

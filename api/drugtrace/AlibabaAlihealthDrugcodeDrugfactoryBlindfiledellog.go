package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugcodeDrugfactoryBlindfiledellog 接收盲底文件删除日志
// alibaba.alihealth.drugcode.drugfactory.blindfiledellog
//
// 临床用药试验-接收盲底文件删除日志
func AlibabaAlihealthDrugcodeDrugfactoryBlindfiledellog(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugcodeDrugfactoryBlindfiledellogAPIRequest, session string) (*drugtrace.AlibabaAlihealthDrugcodeDrugfactoryBlindfiledellogAPIResponse, error) {
	var resp drugtrace.AlibabaAlihealthDrugcodeDrugfactoryBlindfiledellogAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

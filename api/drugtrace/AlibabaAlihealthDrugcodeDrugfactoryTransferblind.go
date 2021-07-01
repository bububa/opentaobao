package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

/* AlibabaAlihealthDrugcodeDrugfactoryTransferblind
传输盲底文件
alibaba.alihealth.drugcode.drugfactory.transferblind

临床用药试验-传输盲底文件 */
func AlibabaAlihealthDrugcodeDrugfactoryTransferblind(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugcodeDrugfactoryTransferblindAPIRequest, session string) (*drugtrace.AlibabaAlihealthDrugcodeDrugfactoryTransferblindAPIResponse, error) {
	var resp drugtrace.AlibabaAlihealthDrugcodeDrugfactoryTransferblindAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

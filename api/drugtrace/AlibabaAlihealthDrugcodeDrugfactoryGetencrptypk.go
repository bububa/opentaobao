package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

/* AlibabaAlihealthDrugcodeDrugfactoryGetencrptypk
获取加密公钥
alibaba.alihealth.drugcode.drugfactory.getencrptypk

获取服务端给药厂用来加密的公钥 */
func AlibabaAlihealthDrugcodeDrugfactoryGetencrptypk(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugcodeDrugfactoryGetencrptypkAPIRequest, session string) (*drugtrace.AlibabaAlihealthDrugcodeDrugfactoryGetencrptypkAPIResponse, error) {
	var resp drugtrace.AlibabaAlihealthDrugcodeDrugfactoryGetencrptypkAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

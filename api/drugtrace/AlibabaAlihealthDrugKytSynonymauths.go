package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytSynonymauths 物流企业查询货主企业信息
// alibaba.alihealth.drug.kyt.synonymauths
//
// 物流企业查询货主企业信息
func AlibabaAlihealthDrugKytSynonymauths(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytSynonymauthsAPIRequest, resp *drugtrace.AlibabaAlihealthDrugKytSynonymauthsAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

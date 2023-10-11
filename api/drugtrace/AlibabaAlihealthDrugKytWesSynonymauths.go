package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytWesSynonymauths 物流企业查询货主企业信息
// alibaba.alihealth.drug.kyt.wes.synonymauths
//
// 物流企业查询货主企业信息
func AlibabaAlihealthDrugKytWesSynonymauths(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytWesSynonymauthsAPIRequest, session string) (*drugtrace.AlibabaAlihealthDrugKytWesSynonymauthsAPIResponse, error) {
	var resp drugtrace.AlibabaAlihealthDrugKytWesSynonymauthsAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

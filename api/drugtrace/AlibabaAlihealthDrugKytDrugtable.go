package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytDrugtable 查询药品目录信息
// alibaba.alihealth.drug.kyt.drugtable
//
// 查询药品目录信息
func AlibabaAlihealthDrugKytDrugtable(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytDrugtableAPIRequest, session string) (*drugtrace.AlibabaAlihealthDrugKytDrugtableAPIResponse, error) {
	var resp drugtrace.AlibabaAlihealthDrugKytDrugtableAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

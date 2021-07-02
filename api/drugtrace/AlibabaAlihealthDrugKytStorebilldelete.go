package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytStorebilldelete 零售端单据删除
// alibaba.alihealth.drug.kyt.storebilldelete
//
// 零售端单据删除
func AlibabaAlihealthDrugKytStorebilldelete(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytStorebilldeleteAPIRequest, session string) (*drugtrace.AlibabaAlihealthDrugKytStorebilldeleteAPIResponse, error) {
	var resp drugtrace.AlibabaAlihealthDrugKytStorebilldeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytQuerydruginfo 码查询药品
// alibaba.alihealth.drug.kyt.querydruginfo
//
// 通过追溯码查询药品信息
func AlibabaAlihealthDrugKytQuerydruginfo(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytQuerydruginfoAPIRequest, session string) (*drugtrace.AlibabaAlihealthDrugKytQuerydruginfoAPIResponse, error) {
	var resp drugtrace.AlibabaAlihealthDrugKytQuerydruginfoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

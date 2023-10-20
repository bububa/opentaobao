package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugCodeKytQuerycode 查询追溯码对应的药品信息
// alibaba.alihealth.drug.code.kyt.querycode
//
// 此接口针对有码药品，提供可通过追溯码获取该药品的基础信息和生产信息；
// 核查平台优先过滤非8开头的，长度非20位数字的码信息。
func AlibabaAlihealthDrugCodeKytQuerycode(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugCodeKytQuerycodeAPIRequest, session string) (*drugtrace.AlibabaAlihealthDrugCodeKytQuerycodeAPIResponse, error) {
	var resp drugtrace.AlibabaAlihealthDrugCodeKytQuerycodeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

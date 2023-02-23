package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugCodeKytWesQuerycode wes查询追溯码信息
// alibaba.alihealth.drug.code.kyt.wes.querycode
//
// 此接口针对有码药品，提供可通过追溯码获取该药品的基础信息和生产信息；
// 核查平台优先过滤非8开头的，长度非20位数字的码信息。
func AlibabaAlihealthDrugCodeKytWesQuerycode(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugCodeKytWesQuerycodeAPIRequest, session string) (*drugtrace.AlibabaAlihealthDrugCodeKytWesQuerycodeAPIResponse, error) {
	var resp drugtrace.AlibabaAlihealthDrugCodeKytWesQuerycodeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

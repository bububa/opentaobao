package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugcodekytyqquerycode 查询追溯码对应的药品信息（疫情）
// alibaba.alihealth.drug.code.kyt.yq.querycode
//
// 通过追溯码码得到 药品名称、包装规格、剂型、剂型规格”、有效期至等信息。
func Alibabaalihealthdrugcodekytyqquerycode(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugcodekytyqquerycodeAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugcodekytyqquerycodeAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugcodekytyqquerycodeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

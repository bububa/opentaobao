package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugcodecommonlistcodeinfo 通用查询码接口
// alibaba.alihealth.drug.code.common.list.codeinfo
//
// 通用查询码接口
func Alibabaalihealthdrugcodecommonlistcodeinfo(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugcodecommonlistcodeinfoAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugcodecommonlistcodeinfoAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugcodecommonlistcodeinfoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

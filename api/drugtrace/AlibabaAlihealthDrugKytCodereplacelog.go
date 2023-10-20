package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugkytcodereplacelog 码替换记录查询
// alibaba.alihealth.drug.kyt.codereplacelog
//
// 码替换记录查询
func Alibabaalihealthdrugkytcodereplacelog(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugkytcodereplacelogAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugkytcodereplacelogAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugkytcodereplacelogAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

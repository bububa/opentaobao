package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugkytcodereplace 单码替换
// alibaba.alihealth.drug.kyt.codereplace
//
// 单码替换
func Alibabaalihealthdrugkytcodereplace(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugkytcodereplaceAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugkytcodereplaceAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugkytcodereplaceAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

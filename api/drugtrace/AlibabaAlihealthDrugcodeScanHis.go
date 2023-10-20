package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugcodescanhis 企业查询扫码历史
// alibaba.alihealth.drugcode.scan.his
//
// 企业查询扫码历史
func Alibabaalihealthdrugcodescanhis(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugcodescanhisAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugcodescanhisAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugcodescanhisAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

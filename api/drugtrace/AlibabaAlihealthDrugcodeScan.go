package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugcodescan 查询扫码信息
// alibaba.alihealth.drugcode.scan
//
// 查询扫码信息
func Alibabaalihealthdrugcodescan(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugcodescanAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugcodescanAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugcodescanAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugcodeerrorreport 码信息错误上报
// alibaba.alihealth.drug.code.error.report
//
// 提供码信息错误上报功能，用于数据校对
func Alibabaalihealthdrugcodeerrorreport(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugcodeerrorreportAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugcodeerrorreportAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugcodeerrorreportAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

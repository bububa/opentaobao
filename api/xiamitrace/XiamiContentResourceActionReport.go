package xiamitrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xiamitrace"
)

// Xiamicontentresourceactionreport 曲库开放平台内容行为上报接口
// xiami.content.resource.action.report
//
// 合作方对接入的曲库开放内容上报行为日志
func Xiamicontentresourceactionreport(clt *core.SDKClient, req *xiamitrace.XiamicontentresourceactionreportAPIRequest, session string) (*xiamitrace.XiamicontentresourceactionreportAPIResponse, error) {
	var resp xiamitrace.XiamicontentresourceactionreportAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

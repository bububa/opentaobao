package util

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/util"
)

// Taobaotopsdkfeedbackupload sdk信息回调
// taobao.top.sdk.feedback.upload
//
// sdk回调客户端基本信息到开放平台，用于做监控之类，有助于帮助isv监控系统稳定性
func Taobaotopsdkfeedbackupload(clt *core.SDKClient, req *util.TaobaotopsdkfeedbackuploadAPIRequest, session string) (*util.TaobaotopsdkfeedbackuploadAPIResponse, error) {
	var resp util.TaobaotopsdkfeedbackuploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

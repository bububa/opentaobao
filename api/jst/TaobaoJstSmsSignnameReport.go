package jst

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jst"
)

// TaobaoJstSmsSignnameReport 服务商存量短信签名上传
// taobao.jst.sms.signname.report
//
// 用于上传目前已经在阿里通信申请到的且正在使用的签名数据，确保签名数据正确，否则会导致短信发送失败！！！
func TaobaoJstSmsSignnameReport(clt *core.SDKClient, req *jst.TaobaoJstSmsSignnameReportAPIRequest, session string) (*jst.TaobaoJstSmsSignnameReportAPIResponse, error) {
	var resp jst.TaobaoJstSmsSignnameReportAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

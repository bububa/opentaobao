package jst

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jst"
)

// TaobaoJstSmsTemplateReport 服务商存量短信模板上传
// taobao.jst.sms.template.report
//
// 用于上传目前已经在阿里通信申请到的且正在使用的模板信息，确保模板数据正确，否则会导致短信发送失败！！！
func TaobaoJstSmsTemplateReport(clt *core.SDKClient, req *jst.TaobaoJstSmsTemplateReportAPIRequest, resp *jst.TaobaoJstSmsTemplateReportAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

package examination

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/examination"
)

// Alibabaalihealthexaminationinvoiceinfonotify 体检机构同步发票信息给阿里健康
// alibaba.alihealth.examination.invoice.info.notify
//
// 体检机构向阿里健康同步发票信息
func Alibabaalihealthexaminationinvoiceinfonotify(clt *core.SDKClient, req *examination.AlibabaalihealthexaminationinvoiceinfonotifyAPIRequest, session string) (*examination.AlibabaalihealthexaminationinvoiceinfonotifyAPIResponse, error) {
	var resp examination.AlibabaalihealthexaminationinvoiceinfonotifyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

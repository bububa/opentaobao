package alihealthcert

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthcert"
)

// Alibabaalihealthexaminationreservecertificatenotify 健康证服务商预约结果通知阿里健康
// alibaba.alihealth.examination.reserve.certificate.notify
//
// 当ISV执行完健康证预约成功之后， 调用通知阿里健康
func Alibabaalihealthexaminationreservecertificatenotify(clt *core.SDKClient, req *alihealthcert.AlibabaalihealthexaminationreservecertificatenotifyAPIRequest, session string) (*alihealthcert.AlibabaalihealthexaminationreservecertificatenotifyAPIResponse, error) {
	var resp alihealthcert.AlibabaalihealthexaminationreservecertificatenotifyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

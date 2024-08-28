package alihealthcert

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthcert"
)

// AlibabaAlihealthExaminationReserveCertificateNotify 健康证服务商预约结果通知阿里健康
// alibaba.alihealth.examination.reserve.certificate.notify
//
// 当ISV执行完健康证预约成功之后， 调用通知阿里健康
func AlibabaAlihealthExaminationReserveCertificateNotify(ctx context.Context, clt *core.SDKClient, req *alihealthcert.AlibabaAlihealthExaminationReserveCertificateNotifyAPIRequest, resp *alihealthcert.AlibabaAlihealthExaminationReserveCertificateNotifyAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

package vaccin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/vaccin"
)

// AlibabaHealthVaccinNoticeOrderCancel 福州疫苗取消预约
// alibaba.health.vaccin.notice.order.cancel
//
// 福州疫苗用户取消预约接口
func AlibabaHealthVaccinNoticeOrderCancel(clt *core.SDKClient, req *vaccin.AlibabaHealthVaccinNoticeOrderCancelAPIRequest, session string) (*vaccin.AlibabaHealthVaccinNoticeOrderCancelAPIResponse, error) {
	var resp vaccin.AlibabaHealthVaccinNoticeOrderCancelAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

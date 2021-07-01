package examination

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/examination"
)

/* AlibabaAlihealthExaminationReserveModifyNotify
通知改期结果
alibaba.alihealth.examination.reserve.modify.notify

体检状态为改期中，服务上通知健康是否改期成功 */
func AlibabaAlihealthExaminationReserveModifyNotify(clt *core.SDKClient, req *examination.AlibabaAlihealthExaminationReserveModifyNotifyAPIRequest, session string) (*examination.AlibabaAlihealthExaminationReserveModifyNotifyAPIResponse, error) {
	var resp examination.AlibabaAlihealthExaminationReserveModifyNotifyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

package examination

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/examination"
)

// Alibabaalihealthexaminationreservemodifynotify 通知改期结果
// alibaba.alihealth.examination.reserve.modify.notify
//
// 体检状态为改期中，服务上通知健康是否改期成功
func Alibabaalihealthexaminationreservemodifynotify(clt *core.SDKClient, req *examination.AlibabaalihealthexaminationreservemodifynotifyAPIRequest, session string) (*examination.AlibabaalihealthexaminationreservemodifynotifyAPIResponse, error) {
	var resp examination.AlibabaalihealthexaminationreservemodifynotifyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

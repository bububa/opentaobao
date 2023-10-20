package examination

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/examination"
)

// Alibabaalihealthexaminationreservecancel 体检机构对接_预约取消
// alibaba.alihealth.examination.reserve.cancel
//
// 体检机构对接_体检取消
func Alibabaalihealthexaminationreservecancel(clt *core.SDKClient, req *examination.AlibabaalihealthexaminationreservecancelAPIRequest, session string) (*examination.AlibabaalihealthexaminationreservecancelAPIResponse, error) {
	var resp examination.AlibabaalihealthexaminationreservecancelAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

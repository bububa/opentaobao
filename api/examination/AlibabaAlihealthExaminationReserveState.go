package examination

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/examination"
)

// Alibabaalihealthexaminationreservestate 体检机构对接_体检状态查询
// alibaba.alihealth.examination.reserve.state
//
// 体检机构对接_体检状态查询
func Alibabaalihealthexaminationreservestate(clt *core.SDKClient, req *examination.AlibabaalihealthexaminationreservestateAPIRequest, session string) (*examination.AlibabaalihealthexaminationreservestateAPIResponse, error) {
	var resp examination.AlibabaalihealthexaminationreservestateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

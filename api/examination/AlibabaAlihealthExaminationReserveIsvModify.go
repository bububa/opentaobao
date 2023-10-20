package examination

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/examination"
)

// Alibabaalihealthexaminationreserveisvmodify ISV调TOP主动发起改期信息
// alibaba.alihealth.examination.reserve.isv.modify
//
// 体检机构对接_ISV发起体检套餐改期
func Alibabaalihealthexaminationreserveisvmodify(clt *core.SDKClient, req *examination.AlibabaalihealthexaminationreserveisvmodifyAPIRequest, session string) (*examination.AlibabaalihealthexaminationreserveisvmodifyAPIResponse, error) {
	var resp examination.AlibabaalihealthexaminationreserveisvmodifyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

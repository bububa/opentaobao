package sungari

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/sungari"
)

// Taobaosungariinspectionsubmit 抽检指令录入
// taobao.sungari.inspection.submit
//
// 抽检指令录入
func Taobaosungariinspectionsubmit(clt *core.SDKClient, req *sungari.TaobaosungariinspectionsubmitAPIRequest, session string) (*sungari.TaobaosungariinspectionsubmitAPIResponse, error) {
	var resp sungari.TaobaosungariinspectionsubmitAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

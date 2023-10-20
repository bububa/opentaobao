package alihealthalgo

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthalgo"
)

// Alibabaalihealthalgomedicationsafetyget 合理用药api
// alibaba.alihealth.algo.medication.safety.get
//
// 合理用药规则引擎服务
func Alibabaalihealthalgomedicationsafetyget(clt *core.SDKClient, req *alihealthalgo.AlibabaalihealthalgomedicationsafetygetAPIRequest, session string) (*alihealthalgo.AlibabaalihealthalgomedicationsafetygetAPIResponse, error) {
	var resp alihealthalgo.AlibabaalihealthalgomedicationsafetygetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

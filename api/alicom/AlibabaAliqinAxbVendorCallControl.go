package alicom

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alicom"
)

// Alibabaaliqinaxbvendorcallcontrol 转呼控制接口
// alibaba.aliqin.axb.vendor.call.control
//
// 转呼控制接口，用于查询小号绑定关系，控制呼叫转接目标
func Alibabaaliqinaxbvendorcallcontrol(clt *core.SDKClient, req *alicom.AlibabaaliqinaxbvendorcallcontrolAPIRequest, session string) (*alicom.AlibabaaliqinaxbvendorcallcontrolAPIResponse, error) {
	var resp alicom.AlibabaaliqinaxbvendorcallcontrolAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

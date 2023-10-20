package alicom

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alicom"
)

// Alibabaaliqinfcvoicegetdetail 获取呼叫详情
// alibaba.aliqin.fc.voice.getdetail
//
// 通过呼叫id获取呼叫相关的数据
func Alibabaaliqinfcvoicegetdetail(clt *core.SDKClient, req *alicom.AlibabaaliqinfcvoicegetdetailAPIRequest, session string) (*alicom.AlibabaaliqinfcvoicegetdetailAPIResponse, error) {
	var resp alicom.AlibabaaliqinfcvoicegetdetailAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

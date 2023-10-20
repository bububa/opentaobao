package alicom

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alicom"
)

// Alibabaaliqinflowpublish 流量发放(用户id)
// alibaba.aliqin.flow.publish
//
// 阿里通信流量下发功能，允许用户补发
func Alibabaaliqinflowpublish(clt *core.SDKClient, req *alicom.AlibabaaliqinflowpublishAPIRequest, session string) (*alicom.AlibabaaliqinflowpublishAPIResponse, error) {
	var resp alicom.AlibabaaliqinflowpublishAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

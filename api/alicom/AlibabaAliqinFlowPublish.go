package alicom

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alicom"
)

/* AlibabaAliqinFlowPublish
流量发放(用户id)
alibaba.aliqin.flow.publish

阿里通信流量下发功能，允许用户补发 */
func AlibabaAliqinFlowPublish(clt *core.SDKClient, req *alicom.AlibabaAliqinFlowPublishAPIRequest, session string) (*alicom.AlibabaAliqinFlowPublishAPIResponse, error) {
	var resp alicom.AlibabaAliqinFlowPublishAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

package aliqin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aliqin"
)

// AlibabaAliqinFcIotQrycard 查询终端信息
// alibaba.aliqin.fc.iot.qrycard
//
// 查询终端信息
func AlibabaAliqinFcIotQrycard(clt *core.SDKClient, req *aliqin.AlibabaAliqinFcIotQrycardAPIRequest, session string) (*aliqin.AlibabaAliqinFcIotQrycardAPIResponse, error) {
	var resp aliqin.AlibabaAliqinFcIotQrycardAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

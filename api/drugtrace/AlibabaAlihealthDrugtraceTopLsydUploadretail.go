package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugtracetoplsyduploadretail 零售单据上传接口
// alibaba.alihealth.drugtrace.top.lsyd.uploadretail
//
// 快易通多融零售上传接口
func Alibabaalihealthdrugtracetoplsyduploadretail(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugtracetoplsyduploadretailAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugtracetoplsyduploadretailAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugtracetoplsyduploadretailAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

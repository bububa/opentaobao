package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugtracetopyljguploadretail 零售单据上传接口
// alibaba.alihealth.drugtrace.top.yljg.uploadretail
//
// 医疗机构零售上传接口
func Alibabaalihealthdrugtracetopyljguploadretail(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugtracetopyljguploadretailAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugtracetopyljguploadretailAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugtracetopyljguploadretailAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

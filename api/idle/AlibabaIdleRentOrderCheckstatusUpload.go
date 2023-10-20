package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// Alibabaidlerentordercheckstatusupload 上传验收结果
// alibaba.idle.rent.order.checkstatus.upload
//
// 上传验收结果
func Alibabaidlerentordercheckstatusupload(clt *core.SDKClient, req *idle.AlibabaidlerentordercheckstatusuploadAPIRequest, session string) (*idle.AlibabaidlerentordercheckstatusuploadAPIResponse, error) {
	var resp idle.AlibabaidlerentordercheckstatusuploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

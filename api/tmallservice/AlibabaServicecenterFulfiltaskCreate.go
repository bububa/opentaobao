package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// Alibabaservicecenterfulfiltaskcreate 合单生成核销单
// alibaba.servicecenter.fulfiltask.create
//
// 服务对工单进行合单，合单的结果是生成核销单
func Alibabaservicecenterfulfiltaskcreate(clt *core.SDKClient, req *tmallservice.AlibabaservicecenterfulfiltaskcreateAPIRequest, session string) (*tmallservice.AlibabaservicecenterfulfiltaskcreateAPIResponse, error) {
	var resp tmallservice.AlibabaservicecenterfulfiltaskcreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

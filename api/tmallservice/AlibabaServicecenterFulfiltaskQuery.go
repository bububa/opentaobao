package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// Alibabaservicecenterfulfiltaskquery 核销单查询
// alibaba.servicecenter.fulfiltask.query
//
// 当系统生成核销单之后，需要派单到服务商，服务商根据核销里的服务信息和用户信息，给消费者提供服务
func Alibabaservicecenterfulfiltaskquery(clt *core.SDKClient, req *tmallservice.AlibabaservicecenterfulfiltaskqueryAPIRequest, session string) (*tmallservice.AlibabaservicecenterfulfiltaskqueryAPIResponse, error) {
	var resp tmallservice.AlibabaservicecenterfulfiltaskqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

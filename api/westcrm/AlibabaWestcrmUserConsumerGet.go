package westcrm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/westcrm"
)

// Alibabawestcrmuserconsumerget 获取指定用户的消费总额
// alibaba.westcrm.user.consumer.get
//
// 获取指定用户的消费总额
func Alibabawestcrmuserconsumerget(clt *core.SDKClient, req *westcrm.AlibabawestcrmuserconsumergetAPIRequest, session string) (*westcrm.AlibabawestcrmuserconsumergetAPIResponse, error) {
	var resp westcrm.AlibabawestcrmuserconsumergetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

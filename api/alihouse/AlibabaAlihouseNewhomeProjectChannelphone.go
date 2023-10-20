package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihousenewhomeprojectchannelphone 新房渠道电话数据同步
// alibaba.alihouse.newhome.project.channelphone
//
// 新房渠道电话数据同步
func Alibabaalihousenewhomeprojectchannelphone(clt *core.SDKClient, req *alihouse.AlibabaalihousenewhomeprojectchannelphoneAPIRequest, session string) (*alihouse.AlibabaalihousenewhomeprojectchannelphoneAPIResponse, error) {
	var resp alihouse.AlibabaalihousenewhomeprojectchannelphoneAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

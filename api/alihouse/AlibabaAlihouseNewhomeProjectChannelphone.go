package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeProjectChannelphone 新房渠道电话数据同步
// alibaba.alihouse.newhome.project.channelphone
//
// 新房渠道电话数据同步
func AlibabaAlihouseNewhomeProjectChannelphone(clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeProjectChannelphoneAPIRequest, session string) (*alihouse.AlibabaAlihouseNewhomeProjectChannelphoneAPIResponse, error) {
	var resp alihouse.AlibabaAlihouseNewhomeProjectChannelphoneAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

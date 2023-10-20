package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabaaxchannelskustatusupdate 翱象渠道商品上下架接口
// alibaba.ax.channel.sku.status.update
//
// 翱象渠道商品上下架接口
func Alibabaaxchannelskustatusupdate(clt *core.SDKClient, req *wdk.AlibabaaxchannelskustatusupdateAPIRequest, session string) (*wdk.AlibabaaxchannelskustatusupdateAPIResponse, error) {
	var resp wdk.AlibabaaxchannelskustatusupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

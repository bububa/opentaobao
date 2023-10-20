package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkumsretrievebatchconfirm 批量消息确认
// alibaba.wdk.ums.retrieve.batch.confirm
//
// 批量消息确认
func Alibabawdkumsretrievebatchconfirm(clt *core.SDKClient, req *wdk.AlibabawdkumsretrievebatchconfirmAPIRequest, session string) (*wdk.AlibabawdkumsretrievebatchconfirmAPIResponse, error) {
	var resp wdk.AlibabawdkumsretrievebatchconfirmAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

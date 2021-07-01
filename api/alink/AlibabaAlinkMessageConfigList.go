package alink

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alink"
)

/* AlibabaAlinkMessageConfigList
查询消息开关配置列表
alibaba.alink.message.config.list

阿里智能获取消息开关配置列表 */
func AlibabaAlinkMessageConfigList(clt *core.SDKClient, req *alink.AlibabaAlinkMessageConfigListAPIRequest, session string) (*alink.AlibabaAlinkMessageConfigListAPIResponse, error) {
	var resp alink.AlibabaAlinkMessageConfigListAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

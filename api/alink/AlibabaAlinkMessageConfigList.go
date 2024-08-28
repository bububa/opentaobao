package alink

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alink"
)

// AlibabaAlinkMessageConfigList 查询消息开关配置列表
// alibaba.alink.message.config.list
//
// 阿里智能获取消息开关配置列表
func AlibabaAlinkMessageConfigList(ctx context.Context, clt *core.SDKClient, req *alink.AlibabaAlinkMessageConfigListAPIRequest, resp *alink.AlibabaAlinkMessageConfigListAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

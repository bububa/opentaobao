package alink

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alink"
)

// AliyunAlinkOpendataUrlQuery 开放数据授权访问URL查询
// aliyun.alink.opendata.url.query
//
// 厂商数据授权访问URL查询
func AliyunAlinkOpendataUrlQuery(ctx context.Context, clt *core.SDKClient, req *alink.AliyunAlinkOpendataUrlQueryAPIRequest, resp *alink.AliyunAlinkOpendataUrlQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

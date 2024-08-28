package campus

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// AlibabaVisitorGetidsbyqrcode 根据访客二维码查访客行程id
// alibaba.visitor.getidsbyqrcode
//
// 根据支付宝阿里访客小程序的动态二维码查询来访行程id
func AlibabaVisitorGetidsbyqrcode(ctx context.Context, clt *core.SDKClient, req *campus.AlibabaVisitorGetidsbyqrcodeAPIRequest, resp *campus.AlibabaVisitorGetidsbyqrcodeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}

package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// AlibabaVisitorGetidsbyqrcode 根据访客二维码查访客行程id
// alibaba.visitor.getidsbyqrcode
//
// 根据支付宝阿里访客小程序的动态二维码查询来访行程id
func AlibabaVisitorGetidsbyqrcode(clt *core.SDKClient, req *campus.AlibabaVisitorGetidsbyqrcodeAPIRequest, session string) (*campus.AlibabaVisitorGetidsbyqrcodeAPIResponse, error) {
	var resp campus.AlibabaVisitorGetidsbyqrcodeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

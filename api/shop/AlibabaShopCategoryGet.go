package shop

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/shop"
)

// Alibabashopcategoryget 指定店铺分类信息查询接口
// alibaba.shop.category.get
//
// 按照卖家身份查询指定分类信息
func Alibabashopcategoryget(clt *core.SDKClient, req *shop.AlibabashopcategorygetAPIRequest, session string) (*shop.AlibabashopcategorygetAPIResponse, error) {
	var resp shop.AlibabashopcategorygetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

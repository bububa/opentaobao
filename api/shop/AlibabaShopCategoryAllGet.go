package shop

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/shop"
)

// Alibabashopcategoryallget 全部店铺分类信息查询接口
// alibaba.shop.category.all.get
//
// 按照卖家身份查询全部分类信息
func Alibabashopcategoryallget(clt *core.SDKClient, req *shop.AlibabashopcategoryallgetAPIRequest, session string) (*shop.AlibabashopcategoryallgetAPIResponse, error) {
	var resp shop.AlibabashopcategoryallgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

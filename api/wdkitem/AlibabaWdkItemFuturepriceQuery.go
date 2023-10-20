package wdkitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdkitem"
)

// Alibabawdkitemfuturepricequery 单个商品未来价查询接口
// alibaba.wdk.item.futureprice.query
//
// 查询单个商品未来价，融合了未来基础售价+未来促销价
func Alibabawdkitemfuturepricequery(clt *core.SDKClient, req *wdkitem.AlibabawdkitemfuturepricequeryAPIRequest, session string) (*wdkitem.AlibabawdkitemfuturepricequeryAPIResponse, error) {
	var resp wdkitem.AlibabawdkitemfuturepricequeryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

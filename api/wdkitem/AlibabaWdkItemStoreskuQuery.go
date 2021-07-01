package wdkitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdkitem"
)

/* AlibabaWdkItemStoreskuQuery
门店商品信息查询
alibaba.wdk.item.storesku.query

门店商品信息查询 */
func AlibabaWdkItemStoreskuQuery(clt *core.SDKClient, req *wdkitem.AlibabaWdkItemStoreskuQueryAPIRequest, session string) (*wdkitem.AlibabaWdkItemStoreskuQueryAPIResponse, error) {
	var resp wdkitem.AlibabaWdkItemStoreskuQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

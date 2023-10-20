package lsticitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/lsticitem"
)

// AlibabaLstIcItemInfoQuery 商品信息查询
// alibaba.lst.ic.item.info.query
//
// 查询商品信息
func AlibabaLstIcItemInfoQuery(clt *core.SDKClient, req *lsticitem.AlibabaLstIcItemInfoQueryAPIRequest, session string) (*lsticitem.AlibabaLstIcItemInfoQueryAPIResponse, error) {
	var resp lsticitem.AlibabaLstIcItemInfoQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

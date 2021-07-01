package icbu

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbu"
)

/* AlibabaIcbuProductList
商品查询
alibaba.icbu.product.list

根据类目ID和商品名称查询商品概要信息。结果以修改时间倒序返回，支持分页，每页最多30个。每次调用都是独立的请求，不记录调用的上下文。 */
func AlibabaIcbuProductList(clt *core.SDKClient, req *icbu.AlibabaIcbuProductListAPIRequest, session string) (*icbu.AlibabaIcbuProductListAPIResponse, error) {
	var resp icbu.AlibabaIcbuProductListAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

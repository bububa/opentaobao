package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// Taobaobaikeimportzhubaodata 导入数据到商品百科服务
// taobao.baike.import.zhubao.data
//
// 用于接入外部数据录入到商品百科中
func Taobaobaikeimportzhubaodata(clt *core.SDKClient, req *product.TaobaobaikeimportzhubaodataAPIRequest, session string) (*product.TaobaobaikeimportzhubaodataAPIResponse, error) {
	var resp product.TaobaobaikeimportzhubaodataAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// Taobaobaikeimportzhubaopicture 百科图片数据导入
// taobao.baike.import.zhubao.picture
//
// 用于接入外部--图片--录入到商品百科中
func Taobaobaikeimportzhubaopicture(clt *core.SDKClient, req *product.TaobaobaikeimportzhubaopictureAPIRequest, session string) (*product.TaobaobaikeimportzhubaopictureAPIResponse, error) {
	var resp product.TaobaobaikeimportzhubaopictureAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

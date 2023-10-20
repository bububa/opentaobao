package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// Tmallproductspecpicupload 上传产品规格认证图片
// tmall.product.spec.pic.upload
//
// 上传指定类型的产品规格认证文件，并返回存有上传成功图片url的产品规格对象
func Tmallproductspecpicupload(clt *core.SDKClient, req *product.TmallproductspecpicuploadAPIRequest, session string) (*product.TmallproductspecpicuploadAPIResponse, error) {
	var resp product.TmallproductspecpicuploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

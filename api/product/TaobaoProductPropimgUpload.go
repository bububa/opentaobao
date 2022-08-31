package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// TaobaoProductPropimgUpload 上传单张产品属性图片，如果需要传多张，可调多次
// taobao.product.propimg.upload
//
// 传入产品ID &lt;br/&gt;传入props,目前仅支持颜色属性.调用taobao.itemprops.get.v2取得颜色属性pid,&lt;br/&gt;再用taobao.itempropvalues.get取得vid;格式:pid:vid,只能传入一个颜色pid:vid串; &lt;br/&gt;传入图片内容 &lt;br/&gt;注意：图片最大为2M,只支持JPG,GIF,如果需要传多张，可调多次
func TaobaoProductPropimgUpload(clt *core.SDKClient, req *product.TaobaoProductPropimgUploadAPIRequest, session string) (*product.TaobaoProductPropimgUploadAPIResponse, error) {
	var resp product.TaobaoProductPropimgUploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

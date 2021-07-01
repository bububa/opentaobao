package omniorder

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/omniorder"
)

/* TaobaoOmniitemItemImageUpload
全渠道商品上传图片
taobao.omniitem.item.image.upload

全渠道商品上传图片 */
func TaobaoOmniitemItemImageUpload(clt *core.SDKClient, req *omniorder.TaobaoOmniitemItemImageUploadAPIRequest, session string) (*omniorder.TaobaoOmniitemItemImageUploadAPIResponse, error) {
	var resp omniorder.TaobaoOmniitemItemImageUploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

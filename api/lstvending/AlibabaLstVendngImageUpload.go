package lstvending

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/lstvending"
)

// Alibabalstvendngimageupload 售货机商品图片上传
// alibaba.lst.vendng.image.upload
//
// 零售通自动售货机商品图片上传接口，主要为ISV厂商提供图片同步的通道，从而建立统一的商品图片库。
func Alibabalstvendngimageupload(clt *core.SDKClient, req *lstvending.AlibabalstvendngimageuploadAPIRequest, session string) (*lstvending.AlibabalstvendngimageuploadAPIResponse, error) {
	var resp lstvending.AlibabalstvendngimageuploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

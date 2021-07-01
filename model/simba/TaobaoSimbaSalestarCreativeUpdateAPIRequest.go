package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSimbaSalestarCreativeUpdateAPIRequest
销量明星更新创意相关接口 API请求
taobao.simba.salestar.creative.update

更新一个创意的信息，可以设置创意标题、创意图片 */
type TaobaoSimbaSalestarCreativeUpdateAPIRequest struct {
	model.Params
	// 推广组Id
	_adgroupId int64
	// 创意Id
	_creativeId int64
	// 创意标题，最多20个汉字
	_title string
	// 创意图片地址，必须是推广组对应商品的图片之一
	_imgUrl string
	// 如果用户开通了创意本地上传图片功能的，可以使用该用户图片空间的图片来修改创意，pictureId为图片空间中图片的pictureId，img_url为图片空间中图片链接地址，如果是使用的主图或副图修改创意，则pictureId必须为空
	_pictureId int64
}

// New

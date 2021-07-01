package viapi

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AliyunViapiGoodstechClassifygoodsAPIRequest
商品分类 API请求
aliyun.viapi.goodstech.classifygoods

可以识别图像中的商品分类，返回商品类目、置信度等信息。目前已经支持服饰鞋包、3C数码、家居用品等超过1万种类目分类。(参数图片/链接必须通过以下方式获取: https://help.aliyun.com/document_detail/155645.html ) */
type AliyunViapiGoodstechClassifygoodsAPIRequest struct {
	model.Params
	// 待检测图片链接
	_imageUrl string
}

// New

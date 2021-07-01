package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSimbaCreativeAddAPIRequest
增加创意 API请求
taobao.simba.creative.add

创建一个创意 */
type TaobaoSimbaCreativeAddAPIRequest struct {
	model.Params
	// 推广组Id
	_adgroupId int64
	// 创意标题，最多20个汉字
	_title string
	// 创意图片地址，必须是推广组对应商品的图片之一
	_imgUrl string
	// 主人昵称
	_nick string
}

// New

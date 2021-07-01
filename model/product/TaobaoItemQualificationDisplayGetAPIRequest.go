package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoItemQualificationDisplayGetAPIRequest
资质采集配置异步获取接口 API请求
taobao.item.qualification.display.get

根据类目，商品，属性等参与动态获得资质采集配置 */
type TaobaoItemQualificationDisplayGetAPIRequest struct {
	model.Params
	// 参数列表，为key和value都是string的map的转化的json格式
	_param string
	// 商品id
	_itemId int64
	// 类目id
	_categoryId int64
}

// New

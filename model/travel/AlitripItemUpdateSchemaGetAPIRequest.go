package travel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripItemUpdateSchemaGetAPIRequest
获取编辑商品的schema模板 API请求
alitrip.item.update.schema.get

获取编辑商品的schema模板。目前支持类目：出境自由行(50278002)、境内自由行(50272002)、出境跟团游(50258005)、境内跟团游(50258004)、境外一日游/多日游(50276003) */
type AlitripItemUpdateSchemaGetAPIRequest struct {
	model.Params
	// 商品id
	_itemId int64
	// 需要获取的编辑schema，不填默认返回全部
	_updateFieldIds []string
}

// New

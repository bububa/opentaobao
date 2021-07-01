package travel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripItemAddSchemaGetAPIRequest
获取商品发布模板 API请求
alitrip.item.add.schema.get

发布飞猪度假商品时，需要先调用此接口获取商品发布的模板schema。目前支持类目：出境自由行(50278002)、境内自由行(50272002)、出境跟团游(50258005)、境内跟团游(50258004)、境外一日游/多日游(50276003) */
type AlitripItemAddSchemaGetAPIRequest struct {
	model.Params
	// 类目id
	_catId int64
}

// New

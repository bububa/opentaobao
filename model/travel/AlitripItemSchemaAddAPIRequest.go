package travel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripItemSchemaAddAPIRequest
使用schema模板进行商品发布 API请求
alitrip.item.schema.add

飞猪度假商品使用schema模板进行商品发布。目前支持类目：出境自由行(50278002)、境内自由行(50272002)、出境跟团游(50258005)、境内跟团游(50258004)、境外一日游/多日游(50276003) */
type AlitripItemSchemaAddAPIRequest struct {
	model.Params
	// 类目id
	_catId int64
	// 商品数据
	_schemaXmlFields string
}

// New

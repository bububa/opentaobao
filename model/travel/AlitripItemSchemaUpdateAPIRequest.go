package travel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripItemSchemaUpdateAPIRequest
使用schema进行商品编辑 API请求
alitrip.item.schema.update

飞猪度假商品使用schema进行商品编辑。目前支持类目：出境自由行(50278002)、境内自由行(50272002)、出境跟团游(50258005)、境内跟团游(50258004)、境外一日游/多日游(50276003) */
type AlitripItemSchemaUpdateAPIRequest struct {
	model.Params
	// 商品id
	_itemId int64
	// 商品数据
	_schemaXmlFields string
}

// New

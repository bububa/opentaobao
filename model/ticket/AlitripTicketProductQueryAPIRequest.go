package ticket

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripTicketProductQueryAPIRequest
【门票API2.0】门票商品查询接口 API请求
alitrip.ticket.product.query

门票商品查询接口：返回商家上传的门票商品信息 */
type AlitripTicketProductQueryAPIRequest struct {
	model.Params
	// 商户自定义收费项目编码。与ali_product_id，item_id 三者至少填写一个
	_outProductId string
	// 阿里标准收费项目id。与out_product_id，item_id 三者至少填写一个
	_aliProductId int64
	// 商品ID。与out_product_id，ali_product_id三者至少填写一个
	_itemId int64
	// 代表业务来源，gongxiao代表供销平台业务
	_pageSource string
}

// New

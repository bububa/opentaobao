package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoPromotionLimitdiscountGetAPIRequest
限时打折查询 API请求
taobao.promotion.limitdiscount.get

分页查询某个卖家的限时打折信息。每页20条数据，按照结束时间降序排列。也可指定某一个限时打折id查询唯一的限时打折信息。 */
type TaobaoPromotionLimitdiscountGetAPIRequest struct {
	model.Params
	// 限时打折ID。这个针对查询唯一限时打折情况。若此字段不为空，则说明操作为单条限时打折记录查询，其他字段忽略。若想分页按条件查询，这个字段置为空。
	_limitDiscountId int64
	// 限时打折活动状态。ALL:全部状态;OVER:已结束;DOING:进行中;PROPARE:未开始(只支持大写)。当limit_discount_id为空时，为空时，默认为全部的状态。
	_status string
	// 限时打折开始时间。输入的时间会被截取，年月日有效，时分秒忽略。
	_startTime string
	// 限时打折结束时间。输入的时间会被截取，年月日有效，时分秒忽略。
	_endTime string
	// 分页页号。默认1。当页数大于最大页数时，结果为最大页数的数据。
	_pageNumber int64
}

// New

package tbk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTbkDgNewuserOrderSumAPIRequest
淘宝客-推广者-拉新活动对应数据查询 API请求
taobao.tbk.dg.newuser.order.sum

拉新活动汇总API */
type TaobaoTbkDgNewuserOrderSumAPIRequest struct {
	model.Params
	// 页大小，默认20，1~100
	_pageSize int64
	// mm_xxx_xxx_xxx的第三位
	_adzoneId int64
	// 页码，默认1
	_pageNo int64
	// mm_xxx_xxx_xxx的第二位
	_siteId int64
	// 活动id， 活动名称与活动ID列表，请参见https://tbk.bbs.taobao.com/detail.html?appId=45301&postId=8599277
	_activityId string
	// 结算月份
	_settleMonth string
}

// New

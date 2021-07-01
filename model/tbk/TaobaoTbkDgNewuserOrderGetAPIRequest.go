package tbk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTbkDgNewuserOrderGetAPIRequest
淘宝客-推广者-新用户订单明细查询 API请求
taobao.tbk.dg.newuser.order.get

拉新API */
type TaobaoTbkDgNewuserOrderGetAPIRequest struct {
	model.Params
	// 页大小，默认20，1~100
	_pageSize int64
	// mm_xxx_xxx_xxx的第三位
	_adzoneId int64
	// 页码，默认1
	_pageNo int64
	// 开始时间，当活动为淘宝活动，表示最早注册时间；当活动为支付宝活动，表示最早绑定时间；当活动为天猫活动，表示最早领取红包时间
	_startTime string
	// 结束时间，当活动为淘宝活动，表示最晚结束时间；当活动为支付宝活动，表示最晚绑定时间；当活动为天猫活动，表示最晚领取红包的时间
	_endTime string
	// 活动id， 活动名称与活动ID列表，请参见https://tbk.bbs.taobao.com/detail.html?appId=45301&postId=8599277
	_activityId string
}

// New

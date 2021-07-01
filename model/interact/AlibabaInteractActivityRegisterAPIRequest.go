package interact

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaInteractActivityRegisterAPIRequest
ISV互动应用活动注册服务 API请求
alibaba.interact.activity.register

为支持卖家由ISV互动应用可以在手淘店铺首页透出，提供ISV互动应用创建的活动注册到手淘的服务 */
type AlibabaInteractActivityRegisterAPIRequest struct {
	model.Params
	// 页面内容链接，仅允许ascii字符
	_entryUrl string
	// 活动ID
	_bizId string
	// 活动描述文字
	_description string
	// 活动结束时间
	_endTime string
	// 活动名称
	_name string
	// 活动封面图片（非必填）
	_picture string
	// 活动开始时间
	_startTime string
	// 是否有有效时间，若为真开始时间和结束时间必填，默认为真
	_hasValidTime bool
}

// New

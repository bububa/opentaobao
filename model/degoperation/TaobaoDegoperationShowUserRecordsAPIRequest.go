package degoperation

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoDegoperationShowUserRecordsAPIRequest
用户中奖记录 API请求
taobao.degoperation.show.user.records

用户中奖记录 */
type TaobaoDegoperationShowUserRecordsAPIRequest struct {
	model.Params
	// 活动后台配置
	_degAppKey string
	// 活动后台配置
	_eventKey string
	// 第几页
	_pageNumber int64
	// 分页尺寸
	_pageSize int64
	// 系统信息
	_degAccessToken string
}

// New

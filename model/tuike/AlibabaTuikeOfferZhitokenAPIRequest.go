package tuike

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaTuikeOfferZhitokenAPIRequest
生成阿里口令 API请求
alibaba.tuike.offer.zhitoken

推荐链接生产吱口令 */
type AlibabaTuikeOfferZhitokenAPIRequest struct {
	model.Params
	// 主标题
	_title string
	// 主要内容
	_desc string
	// 图标链接，可以为空
	_iconLink string
	// SHORT/MEDIUM/LONG
	_bizType string
	// 业务类型
	_source string
	// 生效时间，可以为空
	_startTime int64
	// 左按钮文案，可以为空，默认为"取消"
	_leftBtnText string
	// 左按钮链接，可以为空
	_leftBtnLink string
	// 右按钮文案，可以为空，默认为"确定"
	_rightBtnText string
	// 右按钮链接
	_rightBtnLink string
}

// New

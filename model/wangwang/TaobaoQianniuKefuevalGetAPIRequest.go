package wangwang

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoQianniuKefuevalGetAPIRequest
客服评价详情接口 API请求
taobao.qianniu.kefueval.get

获取买家对客服的服务评价 */
type TaobaoQianniuKefuevalGetAPIRequest struct {
	model.Params
	// 客服的nick，多个用逗号分隔，不要超过10个，带cntaobao的长nick
	_queryIds string
	// 开始时间，格式yyyyMMddHHmmss
	_btime string
	// 结束时间，格式yyyyMMddHHmmss
	_etime string
}

// New

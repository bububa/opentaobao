package servicecenter

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoWeikePerformancePutAPIRequest
提交客服绩效接口 API请求
taobao.weike.performance.put

提交客服绩效接口 */
type TaobaoWeikePerformancePutAPIRequest struct {
	model.Params
	// 订单id
	_id int64
	// 绩效数据封装类
	_perInfoWrapper *PerformanceInfoWrapper
}

// New

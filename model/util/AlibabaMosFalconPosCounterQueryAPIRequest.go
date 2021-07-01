package util

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaMosFalconPosCounterQueryAPIRequest
云POS查看专柜属性 API请求
alibaba.mos.falcon.pos.counter.query

银泰商业获取专柜是否支持小数等属性查看 */
type AlibabaMosFalconPosCounterQueryAPIRequest struct {
	model.Params
	// 设备序列号
	_sn string
	// 门店号
	_storeNo string
	// 专柜号
	_counterNo string
}

// New

package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkBmPaiyangStatDataQueryAPIRequest
派样统计数据查询 API请求
alibaba.wdk.bm.paiyang.stat.data.query

派样统计数据查询 */
type AlibabaWdkBmPaiyangStatDataQueryAPIRequest struct {
	model.Params
	// 入参对象
	_param *PaiyangStatDataParam
}

// New

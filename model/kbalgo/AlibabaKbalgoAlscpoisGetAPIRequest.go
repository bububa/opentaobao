package kbalgo

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaKbalgoAlscpoisGetAPIRequest
百度批量获取本地poi接口 API请求
alibaba.kbalgo.alscpois.get

接口用于百度方获取本地生活poi数据，分页获取。 */
type AlibabaKbalgoAlscpoisGetAPIRequest struct {
	model.Params
	// 第几页
	_pageNum int64
	// 每页的数量。
	_pageSize int64
}

// New

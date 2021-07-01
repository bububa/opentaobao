package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlscCrmOpenPointOperateAPIRequest
积分操作接口 API请求
alibaba.alsc.crm.open.point.operate

同步积分接口 */
type AlibabaAlscCrmOpenPointOperateAPIRequest struct {
	model.Params
	// 入参
	_paramPointOperateOpenReq *PointOperateOpenReq
}

// New

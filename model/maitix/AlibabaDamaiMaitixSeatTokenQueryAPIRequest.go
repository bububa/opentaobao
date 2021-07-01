package maitix

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaDamaiMaitixSeatTokenQueryAPIRequest
分销商选座获取qtoken API请求
alibaba.damai.maitix.seat.token.query

选座分销，分销商查询token */
type AlibabaDamaiMaitixSeatTokenQueryAPIRequest struct {
	model.Params
	// 场次ID-必填
	_performId int64
	// 项目ID-必填
	_projectId int64
	// 必填-选座结束跳转回去的url,这是渠道方自己的url地址,用于接收选座后的座位信息参数
	_callbackUrl string
	// 会话ID，保证一次选座会话,建议使用 appKey+随机串 生成 ；注意：同一个场次下的会话ID不能重复
	_requestId string
}

// New

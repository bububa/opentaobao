package idleisv

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIdleIsvPvListAPIRequest
闲鱼已验货pv查询 API请求
alibaba.idle.isv.pv.list

根据闲鱼渠道类目查询对应的品牌和型号清单，供服务商进行选择 */
type AlibabaIdleIsvPvListAPIRequest struct {
	model.Params
	// 闲鱼渠道类目的id
	_channelCatId string
	// 系统自动生成
	_brandModelInfo []IdleNewPubValueDo
}

// New

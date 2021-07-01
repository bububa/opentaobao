package tttm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AliyunIndustryTttmProductsSyncAPIRequest
天天特卖货品信息同步 API请求
aliyun.industry.tttm.products.sync

天天特卖货品信息同步 */
type AliyunIndustryTttmProductsSyncAPIRequest struct {
	model.Params
	// 产品信息
	_syncProducts []ProductInfoDto
}

// New

package tttm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AliyunIndustryTttmItemsSyncAPIRequest
天天特卖商品信息同步 API请求
aliyun.industry.tttm.items.sync

天天特卖商品信息同步 */
type AliyunIndustryTttmItemsSyncAPIRequest struct {
	model.Params
	// 商品信息
	_syncItems []ItemInfoDto
}

// New

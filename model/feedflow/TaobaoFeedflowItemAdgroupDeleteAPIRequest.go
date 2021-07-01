package feedflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoFeedflowItemAdgroupDeleteAPIRequest
根据单元id删除单元 API请求
taobao.feedflow.item.adgroup.delete

根据单元id删除单元 */
type TaobaoFeedflowItemAdgroupDeleteAPIRequest struct {
	model.Params
	// 计划id
	_campaignId int64
	// 单元id列表
	_adgroupIdList []int64
}

// New

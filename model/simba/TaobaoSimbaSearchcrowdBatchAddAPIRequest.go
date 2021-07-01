package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSimbaSearchcrowdBatchAddAPIRequest
推广单元增加搜索人群 API请求
taobao.simba.searchcrowd.batch.add

推广单元新增搜索人群 */
type TaobaoSimbaSearchcrowdBatchAddAPIRequest struct {
	model.Params
	// 被操作者的淘宝昵称
	_nick string
	// 推广单元id
	_adgroupId int64
	// 新增人群信息,批量接口,入参为list,溢价(discount)范围为[105,400]
	_adgroupTargetingTags string
}

// New

package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSimbaSerchcrowdStateBatchUpdateAPIRequest
单品搜索人群修改状态 API请求
taobao.simba.serchcrowd.state.batch.update

暂停或启用单品推广搜索人群溢价 */
type TaobaoSimbaSerchcrowdStateBatchUpdateAPIRequest struct {
	model.Params
	// 被操作者的淘宝昵称
	_nick string
	// 需要修改出价的人群包id,批量传入时用,分割
	_adgroupCrowdIds []int64
	// 推广单元id
	_adgroupId int64
	// 人群状态,0:暂停;1:启用
	_state int64
}

// New

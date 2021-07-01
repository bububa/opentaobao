package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOmniorderItemTagOperateAPIRequest
全渠道商品打标与去标 API请求
taobao.omniorder.item.tag.operate

用于对全渠道商品进行打标、去标（门店发货标，门店自提标，前置拆单标）操作。另外还包括增加、删除、修改分单系统，接单系统配置。 */
type TaobaoOmniorderItemTagOperateAPIRequest struct {
	model.Params
	// 商品ID
	_itemId int64
	// 商品标,storeDeliver代表门店发货, AllocateByFront代表前置拆单, storeCollect代表门店自提
	_types []string
	// 操作状态， 填 1 代表打标，填 -1 代表去标
	_status int64
	// 分单&接单设置
	_omniSetting *OmniSettingDto
}

// New

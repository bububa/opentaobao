package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoInventoryStoreManageAPIRequest
创建或更新仓库 API请求
taobao.inventory.store.manage

创建商家仓或者更新商家仓信息 */
type TaobaoInventoryStoreManageAPIRequest struct {
	model.Params
	// 参数定义，ADD：新建; UPDATE：更新
	_operateType string
	// 商家的仓库编码，不允许重复，不允许更新
	_storeCode string
	// 商家的仓库名称，可更新
	_storeName string
	// 仓库类型，可更新。目前只支持自有仓，TYPE_OWN：自有物理仓
	_storeType string
	// 仓库简称，可更新
	_aliasName string
	// 仓库的物理地址，可更新
	_address string
	// 仓库区域名，可更新
	_addressAreaName string
	// 联系人，可更新
	_contact string
	// 联系电话，可更新
	_phone string
	// 邮编，可更新
	_postcode int64
}

// New

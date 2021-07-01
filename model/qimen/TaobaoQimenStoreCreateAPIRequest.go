package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoQimenStoreCreateAPIRequest
门店新增接口 API请求
taobao.qimen.store.create

isv调用接口来讲线下门店同步到线上 */
type TaobaoQimenStoreCreateAPIRequest struct {
	model.Params
	// 门店名称
	_storeName string
	// 门店主营类目
	_mainCategory int64
	// 商户名称
	_companyName string
	// 关闭营业时间(只填时，分；只支持半点和整点)
	_endTime string
	// 开始营业时间(只填时，分；只支持半点和整点)
	_startTime string
	// 门店状态
	_storeStatus string
	// 商户介绍
	_storeDescription string
	// 地址信息
	_address *Address
	// 需要关联的线上店铺ID
	_shopId int64
	// 门店所有者信息
	_storeKeeper *StoreKeeper
	// 门店的类型
	_storeType string
	// ERP系统中门店的编码
	_storeCode string
	// 备注
	_remark string
}

// New

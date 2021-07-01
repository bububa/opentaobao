package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoQimenStoreUpdateAPIRequest
门店更新接口 API请求
taobao.qimen.store.update

商家在ERP等系统中调用该接口，更新门店信息 */
type TaobaoQimenStoreUpdateAPIRequest struct {
	model.Params
	// 门店名称
	_storeName string
	// 备注
	_remark string
	// 门店主营类目
	_mainCategory int64
	// 停止营业时间(只填时，分；只支持半点和整点)
	_endTime string
	// 商户名称
	_companyName string
	// 开始营业时间(只填时，分；只支持半点和整点)
	_startTime string
	// 门店状态
	_storeStatus string
	// 商户介绍
	_storeDescription string
	// 门店地址信息
	_address *Address
	// 需要关联的线上店铺ID
	_shopId int64
	// 门店所有者信息
	_storeKeeper *StoreKeeper
	// 门店类型
	_storeType string
	// 线上门店id
	_storeId int64
	// ERP系统中 门店编码
	_storeCode string
}

// New

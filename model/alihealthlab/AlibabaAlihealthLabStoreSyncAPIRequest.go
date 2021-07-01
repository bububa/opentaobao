package alihealthlab

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthLabStoreSyncAPIRequest
阿里健康检验检测业务，isv门店同步到健康 API请求
alibaba.alihealth.lab.store.sync

阿里健康检验检测业务，isv门店同步到健康。支持门店的上线、下线操作 */
type AlibabaAlihealthLabStoreSyncAPIRequest struct {
	model.Params
	// EFFECTIVE 生效，INVALID 失效
	_isvStoreStatus string
	// 预约须知
	_reserveNotice string
	// 支持在线报告
	_supportOnlineReport bool
	// 门店类型描述
	_storeTypeDesc string
	// 企业社会征信号
	_socialCreditCode string
	// 营业执照编号
	_licenseNo string
	// 营业执照名称
	_licenseName string
	// 门店交通路线
	_storeRoutesDesc string
	// 营业时间描述
	_workTimeDesc string
	// 门店电话号码
	_storePhone string
	// 门店介绍
	_storeIntro string
	// 经度
	_longitude *BigDecimal
	// 纬度
	_latitude *BigDecimal
	// 城市编码
	_cityCode int64
	// 门店地址
	_storeAddress string
	// isv门店编码
	_isvStoreCode string
	// 门店名称
	_storeName string
	// 支持的淘宝商品类目ID，阿里医院场景
	_allowedTbItemCategoryIds []int64
}

// New

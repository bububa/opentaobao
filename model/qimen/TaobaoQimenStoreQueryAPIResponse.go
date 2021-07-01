package qimen

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoQimenStoreQueryAPIResponse
门店信息查询接口 API返回值
taobao.qimen.store.query

商家在ERP等系统中调用该接口，查询门店相关信息 */
type TaobaoQimenStoreQueryAPIResponse struct {
	model.CommonResponse
	TaobaoQimenStoreQueryAPIResponseModel
}

// TaobaoQimenStoreQueryAPIResponseModel is 门店信息查询接口 成功返回结果
type TaobaoQimenStoreQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"qimen_store_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 门店名称
	StoreName string `json:"store_name,omitempty" xml:"store_name,omitempty"`
	// 备注
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 门店主营类目
	MainCategory int64 `json:"main_category,omitempty" xml:"main_category,omitempty"`
	// 响应code
	QimenCode string `json:"qimen_code,omitempty" xml:"qimen_code,omitempty"`
	// 关闭营业时间
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 商户名称
	CompanyName string `json:"company_name,omitempty" xml:"company_name,omitempty"`
	// 开始营业时间
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 响应消息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 门店状态
	StoreStatus string `json:"store_status,omitempty" xml:"store_status,omitempty"`
	// 响应标示
	Flag string `json:"flag,omitempty" xml:"flag,omitempty"`
	// 商户介绍
	StoreDescription string `json:"store_description,omitempty" xml:"store_description,omitempty"`
	// 地址信息
	Address *Address `json:"address,omitempty" xml:"address,omitempty"`
	// 需要关联的线上店铺ID
	ShopId string `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
	// 所有者信息
	StoreKeeper *StoreKeeper `json:"store_keeper,omitempty" xml:"store_keeper,omitempty"`
	// 类型
	StoreType string `json:"store_type,omitempty" xml:"store_type,omitempty"`
	// ERP系统中 门店编码
	StoreCode string `json:"store_code,omitempty" xml:"store_code,omitempty"`
}

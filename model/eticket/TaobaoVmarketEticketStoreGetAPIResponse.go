package eticket

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoVmarketEticketStoreGetAPIResponse 获取电子凭证预约门店信息 API返回值
// taobao.vmarket.eticket.store.get
//
// 用于给外部商家查询电子凭证预约门店信息
type TaobaoVmarketEticketStoreGetAPIResponse struct {
	model.CommonResponse
	TaobaoVmarketEticketStoreGetAPIResponseModel
}

// TaobaoVmarketEticketStoreGetAPIResponseModel is 获取电子凭证预约门店信息 成功返回结果
type TaobaoVmarketEticketStoreGetAPIResponseModel struct {
	XMLName xml.Name `xml:"vmarket_eticket_store_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 商户地址
	Address string `json:"address,omitempty" xml:"address,omitempty"`
	// 商户名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 所在城市
	City string `json:"city,omitempty" xml:"city,omitempty"`
	// 省份
	Province string `json:"province,omitempty" xml:"province,omitempty"`
	// 区
	District string `json:"district,omitempty" xml:"district,omitempty"`
	// 联系电话
	Contract string `json:"contract,omitempty" xml:"contract,omitempty"`
	// 自有卖家导入门店的时候，可以把自己系统门店信息的主键或者唯一key传入，用于快速匹配
	Selfcode string `json:"selfcode,omitempty" xml:"selfcode,omitempty"`
	// 商户id
	StoreId int64 `json:"store_id,omitempty" xml:"store_id,omitempty"`
}

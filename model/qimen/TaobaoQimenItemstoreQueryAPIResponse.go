package qimen

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoQimenItemstoreQueryAPIResponse
商品关联门店查询接口 API返回值
taobao.qimen.itemstore.query

商家在ERP等系统中调用该接口，查询线上商品所关联的门店列表 */
type TaobaoQimenItemstoreQueryAPIResponse struct {
	model.CommonResponse
	TaobaoQimenItemstoreQueryAPIResponseModel
}

// TaobaoQimenItemstoreQueryAPIResponseModel is 商品关联门店查询接口 成功返回结果
type TaobaoQimenItemstoreQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"qimen_itemstore_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 门店列表
	StoreIds []int64 `json:"store_ids,omitempty" xml:"store_ids>int64,omitempty"`
	// 响应的标签
	Flag string `json:"flag,omitempty" xml:"flag,omitempty"`
	// 总的门店数
	TotalLines int64 `json:"total_lines,omitempty" xml:"total_lines,omitempty"`
	// 响应的code
	QimenCode string `json:"qimen_code,omitempty" xml:"qimen_code,omitempty"`
}

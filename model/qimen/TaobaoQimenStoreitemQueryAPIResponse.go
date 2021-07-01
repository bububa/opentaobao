package qimen

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoQimenStoreitemQueryAPIResponse
门店关联商品查询接口 API返回值
taobao.qimen.storeitem.query

商家在ERP等系统中调用该接口，查询某门店所关联的线上商品列表 */
type TaobaoQimenStoreitemQueryAPIResponse struct {
	model.CommonResponse
	TaobaoQimenStoreitemQueryAPIResponseModel
}

// TaobaoQimenStoreitemQueryAPIResponseModel is 门店关联商品查询接口 成功返回结果
type TaobaoQimenStoreitemQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"qimen_storeitem_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应消息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 商品列表
	ItemIds []string `json:"item_ids,omitempty" xml:"item_ids>string,omitempty"`
	// 响应标签
	Flag string `json:"flag,omitempty" xml:"flag,omitempty"`
	// 商品总数
	TotalLines int64 `json:"total_lines,omitempty" xml:"total_lines,omitempty"`
	// 响应code
	QimenCode string `json:"qimen_code,omitempty" xml:"qimen_code,omitempty"`
}

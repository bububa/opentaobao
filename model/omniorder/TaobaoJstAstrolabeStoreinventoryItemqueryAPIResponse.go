package omniorder

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaojstastrolabestoreinventoryitemqueryAPIResponse 库存查询接口 API返回值
// taobao.jst.astrolabe.storeinventory.itemquery
//
// 查询门店或电商仓库存，该接口一次可以同时查询多个门店或电商仓的多个商品的多种类型的库存
type TaobaojstastrolabestoreinventoryitemqueryAPIResponse struct {
	model.CommonResponse
	TaobaojstastrolabestoreinventoryitemqueryAPIResponseModel
}

// TaobaojstastrolabestoreinventoryitemqueryAPIResponseModel is 库存查询接口 成功返回结果
type TaobaojstastrolabestoreinventoryitemqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"jst_astrolabe_storeinventory_itemquery_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 门店列表
	Stores []Store `json:"stores,omitempty" xml:"stores>store,omitempty"`
	// 响应标示
	Flag string `json:"flag,omitempty" xml:"flag,omitempty"`
	// 响应编码
	QimenCode string `json:"qimen_code,omitempty" xml:"qimen_code,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
}

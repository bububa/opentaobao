package qimen

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenStorecategoryGetAPIResponse 门店类目获取接口 API返回值
// taobao.qimen.storecategory.get
//
// 商家在ERP中调用该接口，获取门店类目
type TaobaoQimenStorecategoryGetAPIResponse struct {
	model.CommonResponse
	TaobaoQimenStorecategoryGetAPIResponseModel
}

// TaobaoQimenStorecategoryGetAPIResponseModel is 门店类目获取接口 成功返回结果
type TaobaoQimenStorecategoryGetAPIResponseModel struct {
	XMLName xml.Name `xml:"qimen_storecategory_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 响应标示
	Flag string `json:"flag,omitempty" xml:"flag,omitempty"`
	// 响应code
	QimenCode string `json:"qimen_code,omitempty" xml:"qimen_code,omitempty"`
	// 类目json字符串
	StoreCategory string `json:"store_category,omitempty" xml:"store_category,omitempty"`
}

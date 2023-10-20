package fivee

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaofiveeinnerproductgetAPIResponse 国产商品资质查询 API返回值
// taobao.fivee.innerproduct.get
//
// 资质共享平台，国产商品查询
type TaobaofiveeinnerproductgetAPIResponse struct {
	model.CommonResponse
	TaobaofiveeinnerproductgetAPIResponseModel
}

// TaobaofiveeinnerproductgetAPIResponseModel is 国产商品资质查询 成功返回结果
type TaobaofiveeinnerproductgetAPIResponseModel struct {
	XMLName xml.Name `xml:"fivee_innerproduct_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TaobaofiveeinnerproductgetResult `json:"result,omitempty" xml:"result,omitempty"`
}

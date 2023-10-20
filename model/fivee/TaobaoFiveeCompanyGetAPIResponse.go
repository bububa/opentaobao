package fivee

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaofiveecompanygetAPIResponse 查询商信息 API返回值
// taobao.fivee.company.get
//
// 资质共享平台查询商信息
type TaobaofiveecompanygetAPIResponse struct {
	model.CommonResponse
	TaobaofiveecompanygetAPIResponseModel
}

// TaobaofiveecompanygetAPIResponseModel is 查询商信息 成功返回结果
type TaobaofiveecompanygetAPIResponseModel struct {
	XMLName xml.Name `xml:"fivee_company_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TaobaofiveecompanygetResult `json:"result,omitempty" xml:"result,omitempty"`
}

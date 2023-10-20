package tbk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTbkScMembergroupOptionalAPIResponse 工具服务商member组查询、新增接口 API返回值
// taobao.tbk.sc.membergroup.optional
//
// 工具服务商member组查询、新增接口
type TaobaoTbkScMembergroupOptionalAPIResponse struct {
	model.CommonResponse
	TaobaoTbkScMembergroupOptionalAPIResponseModel
}

// TaobaoTbkScMembergroupOptionalAPIResponseModel is 工具服务商member组查询、新增接口 成功返回结果
type TaobaoTbkScMembergroupOptionalAPIResponseModel struct {
	XMLName xml.Name `xml:"tbk_sc_membergroup_optional_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Data *TaobaoTbkScMembergroupOptionalMapData `json:"data,omitempty" xml:"data,omitempty"`
}

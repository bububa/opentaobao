package tbk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaotbkscmembergroupoptionalAPIResponse 工具服务商member组查询、新增接口 API返回值
// taobao.tbk.sc.membergroup.optional
//
// 工具服务商member组查询、新增接口
type TaobaotbkscmembergroupoptionalAPIResponse struct {
	model.CommonResponse
	TaobaotbkscmembergroupoptionalAPIResponseModel
}

// TaobaotbkscmembergroupoptionalAPIResponseModel is 工具服务商member组查询、新增接口 成功返回结果
type TaobaotbkscmembergroupoptionalAPIResponseModel struct {
	XMLName xml.Name `xml:"tbk_sc_membergroup_optional_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Data *TaobaotbkscmembergroupoptionalMapData `json:"data,omitempty" xml:"data,omitempty"`
}

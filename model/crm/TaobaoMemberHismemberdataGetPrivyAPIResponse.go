package crm

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaomemberhismemberdatagetprivyAPIResponse 会员历史备份数据查询 API返回值
// taobao.member.hismemberdata.get.privy
//
// 会员历史备份数据分页查询，查询内容为等级，会员备份版本，会员nick等信息.
type TaobaomemberhismemberdatagetprivyAPIResponse struct {
	model.CommonResponse
	TaobaomemberhismemberdatagetprivyAPIResponseModel
}

// TaobaomemberhismemberdatagetprivyAPIResponseModel is 会员历史备份数据查询 成功返回结果
type TaobaomemberhismemberdatagetprivyAPIResponseModel struct {
	XMLName xml.Name `xml:"member_hismemberdata_get_privy_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果对象
	ResultDto *TaobaomemberhismemberdatagetprivyResultDto `json:"result_dto,omitempty" xml:"result_dto,omitempty"`
}

package crm

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaocrmgrademktmemberdetailcreateAPIResponse 会员等级营销-创建商品等级营销明细 API返回值
// taobao.crm.grademkt.member.detail.create
//
// 创建商品等级营销明细
type TaobaocrmgrademktmemberdetailcreateAPIResponse struct {
	model.CommonResponse
	TaobaocrmgrademktmemberdetailcreateAPIResponseModel
}

// TaobaocrmgrademktmemberdetailcreateAPIResponseModel is 会员等级营销-创建商品等级营销明细 成功返回结果
type TaobaocrmgrademktmemberdetailcreateAPIResponseModel struct {
	XMLName xml.Name `xml:"crm_grademkt_member_detail_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// json格式
	Module bool `json:"module,omitempty" xml:"module,omitempty"`
}

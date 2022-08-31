package alscmerchant

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscMerchantExtTicketQueryAPIResponse 口碑凭证模版查询服务 API返回值
// alibaba.alsc.merchant.ext.ticket.query
//
// isv需要在c端展示凭证信息,需要查询凭证模版
type AlibabaAlscMerchantExtTicketQueryAPIResponse struct {
	model.CommonResponse
	AlibabaAlscMerchantExtTicketQueryAPIResponseModel
}

// AlibabaAlscMerchantExtTicketQueryAPIResponseModel is 口碑凭证模版查询服务 成功返回结果
type AlibabaAlscMerchantExtTicketQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alsc_merchant_ext_ticket_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 请求requestID
	TicketRequestId string `json:"ticket_request_id,omitempty" xml:"ticket_request_id,omitempty"`
	// 业务编码
	BizCode string `json:"biz_code,omitempty" xml:"biz_code,omitempty"`
	// 外部AppID(外部券)
	OutAppId string `json:"out_app_id,omitempty" xml:"out_app_id,omitempty"`
	// 核销规则(方式) USER_PAY_CODE(支核一体) TICKET_CODE(券码核销) EXTERNAL_TICKET_CODE(外部券核销)
	TicketUseRule string `json:"ticket_use_rule,omitempty" xml:"ticket_use_rule,omitempty"`
	// 凭证图片(完整地址)
	TicketTemplateImg string `json:"ticket_template_img,omitempty" xml:"ticket_template_img,omitempty"`
	// 品牌logo(完整地址)
	BrandLogo string `json:"brand_logo,omitempty" xml:"brand_logo,omitempty"`
	// 凭证模版品牌名
	BrandName string `json:"brand_name,omitempty" xml:"brand_name,omitempty"`
	// 凭证模版业务类型(&#34;SERV_INDUSTRY&#34;,&#34;生服&#34;,&#34;CATERING&#34;,&#34;餐饮&#34;,&#34;MALL&#34;&#34;商圈&#34;)
	TicketTemplateType string `json:"ticket_template_type,omitempty" xml:"ticket_template_type,omitempty"`
	// 凭证模版业务名称
	TicketTemplateName string `json:"ticket_template_name,omitempty" xml:"ticket_template_name,omitempty"`
	// 凭证模版ID
	TicketTemplateId string `json:"ticket_template_id,omitempty" xml:"ticket_template_id,omitempty"`
	// 备注
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 外部商品ID(多个用,分割)(外部券)
	OutItemId string `json:"out_item_id,omitempty" xml:"out_item_id,omitempty"`
	// 使用须知
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// 现价(售价)(单位分)
	CurrentPrice int64 `json:"current_price,omitempty" xml:"current_price,omitempty"`
	// 原价(单位分)
	OriginalPrice int64 `json:"original_price,omitempty" xml:"original_price,omitempty"`
	// 有效期
	TicketTemplateValidPeriod *TicketTemplateValidPeriod `json:"ticket_template_valid_period,omitempty" xml:"ticket_template_valid_period,omitempty"`
	// 单次可使用数量(支核一体用)
	MaxVerifyCountToOrder int64 `json:"max_verify_count_to_order,omitempty" xml:"max_verify_count_to_order,omitempty"`
}

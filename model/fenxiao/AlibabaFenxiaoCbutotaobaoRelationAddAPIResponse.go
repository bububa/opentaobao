package fenxiao

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaFenxiaoCbutotaobaoRelationAddAPIResponse 1688分销铺货到淘宝关系添加 API返回值
// alibaba.fenxiao.cbutotaobao.relation.add
//
// 1688分销铺货到淘宝关系添加
type AlibabaFenxiaoCbutotaobaoRelationAddAPIResponse struct {
	model.CommonResponse
	AlibabaFenxiaoCbutotaobaoRelationAddAPIResponseModel
}

// AlibabaFenxiaoCbutotaobaoRelationAddAPIResponseModel is 1688分销铺货到淘宝关系添加 成功返回结果
type AlibabaFenxiaoCbutotaobaoRelationAddAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_fenxiao_cbutotaobao_relation_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果返回值
	Result string `json:"result,omitempty" xml:"result,omitempty"`
	// 错误信息
	ResultMessage string `json:"result_message,omitempty" xml:"result_message,omitempty"`
	// 是否成功
	ResultSuccess bool `json:"result_success,omitempty" xml:"result_success,omitempty"`
}

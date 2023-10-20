package nrt

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallnrtcoupontemplatesynAPIResponse 喵零券同步 API返回值
// tmall.nrt.coupon.template.syn
//
// 查询券模版
type TmallnrtcoupontemplatesynAPIResponse struct {
	model.CommonResponse
	TmallnrtcoupontemplatesynAPIResponseModel
}

// TmallnrtcoupontemplatesynAPIResponseModel is 喵零券同步 成功返回结果
type TmallnrtcoupontemplatesynAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_nrt_coupon_template_syn_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TmallnrtcoupontemplatesynResult `json:"result,omitempty" xml:"result,omitempty"`
}

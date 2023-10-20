package alilabs

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoailabaicloudtopskilslistnewAPIResponse 获取产品下挂载的技能列表 API返回值
// taobao.ailab.aicloud.top.skils.list.new
//
// 星空平台提供的获取产品下挂载的技能列表新接口
type TaobaoailabaicloudtopskilslistnewAPIResponse struct {
	model.CommonResponse
	TaobaoailabaicloudtopskilslistnewAPIResponseModel
}

// TaobaoailabaicloudtopskilslistnewAPIResponseModel is 获取产品下挂载的技能列表 成功返回结果
type TaobaoailabaicloudtopskilslistnewAPIResponseModel struct {
	XMLName xml.Name `xml:"ailab_aicloud_top_skils_list_new_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回包装类
	Result *BaseResult `json:"result,omitempty" xml:"result,omitempty"`
}

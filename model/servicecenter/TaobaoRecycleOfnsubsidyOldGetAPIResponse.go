package servicecenter

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoRecycleOfnsubsidyOldGetAPIResponse 回收单旧机款及补贴查询 API返回值
// taobao.recycle.ofnsubsidy.old.get
//
// 回收单旧机款及补贴查询
type TaobaoRecycleOfnsubsidyOldGetAPIResponse struct {
	model.CommonResponse
	TaobaoRecycleOfnsubsidyOldGetAPIResponseModel
}

// TaobaoRecycleOfnsubsidyOldGetAPIResponseModel is 回收单旧机款及补贴查询 成功返回结果
type TaobaoRecycleOfnsubsidyOldGetAPIResponseModel struct {
	XMLName xml.Name `xml:"recycle_ofnsubsidy_old_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 回收款及补贴
	Data *OfnRecycleOrderSubsidyDto `json:"data,omitempty" xml:"data,omitempty"`
}

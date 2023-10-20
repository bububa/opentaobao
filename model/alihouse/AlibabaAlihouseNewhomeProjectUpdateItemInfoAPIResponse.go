package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousenewhomeprojectupdateiteminfoAPIResponse 更新楼盘商品信息 API返回值
// alibaba.alihouse.newhome.project.update.item.info
//
// 更新楼盘商品信息
type AlibabaalihousenewhomeprojectupdateiteminfoAPIResponse struct {
	model.CommonResponse
	AlibabaalihousenewhomeprojectupdateiteminfoAPIResponseModel
}

// AlibabaalihousenewhomeprojectupdateiteminfoAPIResponseModel is 更新楼盘商品信息 成功返回结果
type AlibabaalihousenewhomeprojectupdateiteminfoAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_project_update_item_info_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 实体类
	Result *AlibabaalihousenewhomeprojectupdateiteminfoResult `json:"result,omitempty" xml:"result,omitempty"`
}

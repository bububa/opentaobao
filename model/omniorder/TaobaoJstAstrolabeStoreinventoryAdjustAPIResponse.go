package omniorder

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaojstastrolabestoreinventoryadjustAPIResponse 后端商品库存占用调整接口 API返回值
// taobao.jst.astrolabe.storeinventory.adjust
//
// 当第三方系统出现分单结果和天猫货品中心分单结果不一致时，需要调用此接口同步分单消息给天猫货品中心，调整之前占用的门店/电商仓库存。
type TaobaojstastrolabestoreinventoryadjustAPIResponse struct {
	model.CommonResponse
	TaobaojstastrolabestoreinventoryadjustAPIResponseModel
}

// TaobaojstastrolabestoreinventoryadjustAPIResponseModel is 后端商品库存占用调整接口 成功返回结果
type TaobaojstastrolabestoreinventoryadjustAPIResponseModel struct {
	XMLName xml.Name `xml:"jst_astrolabe_storeinventory_adjust_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应标签
	Flag string `json:"flag,omitempty" xml:"flag,omitempty"`
	// 响应编码
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 响应信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
}

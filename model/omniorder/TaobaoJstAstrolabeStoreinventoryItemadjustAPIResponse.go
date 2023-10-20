package omniorder

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJstAstrolabeStoreinventoryItemadjustAPIResponse 库存占用调整接口 API返回值
// taobao.jst.astrolabe.storeinventory.itemadjust
//
// 当第三方系统出现分单结果和天猫货品中心分单结果不一致时，需要调用此接口同步分单消息给天猫货品中心，调整之前占用的门店/电商仓库存。
type TaobaoJstAstrolabeStoreinventoryItemadjustAPIResponse struct {
	model.CommonResponse
	TaobaoJstAstrolabeStoreinventoryItemadjustAPIResponseModel
}

// TaobaoJstAstrolabeStoreinventoryItemadjustAPIResponseModel is 库存占用调整接口 成功返回结果
type TaobaoJstAstrolabeStoreinventoryItemadjustAPIResponseModel struct {
	XMLName xml.Name `xml:"jst_astrolabe_storeinventory_itemadjust_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应标签
	Flag string `json:"flag,omitempty" xml:"flag,omitempty"`
	// 响应编码
	QimenCode string `json:"qimen_code,omitempty" xml:"qimen_code,omitempty"`
	// 响应信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
}

package smartstore

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSmartstoreDeviceIashelfFeedbackAPIResponse 智能硬件云货架数据回流 API返回值
// taobao.smartstore.device.iashelf.feedback
//
// 智慧门店云货架设备回流规则：（互动云货架、VR云货架通用）</br>
// 1.回流的数据属于当前授权的用户，回流的设备device_code由当前应用添加</br>
// 2.对于快闪店的智能硬件不需要授权</br>
// 3.每一个action都必须传入用户操作时间op_time（start/end_time后续废弃）</br>
// 4.action为ITEM_CLICK、ITEM_SENSOR、BUY_CLICK时，item_id必须传入,且必须是淘宝商品的数字id</br>
// 5.outer_biz_id 用于硬件设备大量数据回流场景，服务商本地日志统计系统对一条日志记录生成唯一标识。 平台后端会对传入的outer_biz_id 做去重处理</br>
// 6.outer_user 用于标识不能获取淘宝账号的游客</br>
type TaobaoSmartstoreDeviceIashelfFeedbackAPIResponse struct {
	model.CommonResponse
	TaobaoSmartstoreDeviceIashelfFeedbackAPIResponseModel
}

// TaobaoSmartstoreDeviceIashelfFeedbackAPIResponseModel is 智能硬件云货架数据回流 成功返回结果
type TaobaoSmartstoreDeviceIashelfFeedbackAPIResponseModel struct {
	XMLName xml.Name `xml:"smartstore_device_iashelf_feedback_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 回流结果
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}

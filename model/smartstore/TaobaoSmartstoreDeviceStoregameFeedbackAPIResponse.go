package smartstore

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
智能硬件互动大屏数据回流 API返回值 
taobao.smartstore.device.storegame.feedback

智慧门店互动引流屏设备回流规则：（适用于智慧迎宾屏、互动游戏、互动拍照、娃娃机等）</br>
1.回流的数据属于当前授权的用户，回流的设备device_code由当前应用添加</br>
2.对于快闪店的智能硬件不需要授权</br>
3.每一个action都必须传入用户操作时间op_time；（start/end_time后续废弃）</br>
4.action为WINNING_PRIZE时，需传入draw_result，只能传入0或者1</br>
5.outer_biz_id 用于硬件设备大量数据回流场景，服务商本地日志统计系统对一条日志记录生成唯一标识。 平台后端会对传入的outer_biz_id 做去重处理</br>
6.outer_user 用于标识不能获取淘宝账号的游客</br>
*/
type TaobaoSmartstoreDeviceStoregameFeedbackAPIResponse struct {
    model.CommonResponse
    TaobaoSmartstoreDeviceStoregameFeedbackAPIResponseModel
}

// 智能硬件互动大屏数据回流 成功返回结果
type TaobaoSmartstoreDeviceStoregameFeedbackAPIResponseModel struct {
    XMLName xml.Name `xml:"smartstore_device_storegame_feedback_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 回流结果
    Result   string `json:"result,omitempty" xml:"result,omitempty"`
}

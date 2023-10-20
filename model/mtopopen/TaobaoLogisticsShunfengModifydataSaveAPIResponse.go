package mtopopen

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaologisticsshunfengmodifydatasaveAPIResponse 顺丰小程序修改配送信息回传接口 API返回值
// taobao.logistics.shunfeng.modifydata.save
//
// 顺丰小程序修改配送信息回传接口
type TaobaologisticsshunfengmodifydatasaveAPIResponse struct {
	model.CommonResponse
	TaobaologisticsshunfengmodifydatasaveAPIResponseModel
}

// TaobaologisticsshunfengmodifydatasaveAPIResponseModel is 顺丰小程序修改配送信息回传接口 成功返回结果
type TaobaologisticsshunfengmodifydatasaveAPIResponseModel struct {
	XMLName xml.Name `xml:"logistics_shunfeng_modifydata_save_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 通讯失败码
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 通讯失败信息
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 回传返回结果对象
	ModifyResponse *ModifyResponse `json:"modify_response,omitempty" xml:"modify_response,omitempty"`
	// 通讯成功/失败
	ResultSuccess bool `json:"result_success,omitempty" xml:"result_success,omitempty"`
}

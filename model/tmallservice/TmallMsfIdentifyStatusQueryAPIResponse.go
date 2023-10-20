package tmallservice

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallMsfIdentifyStatusQueryAPIResponse 喵师傅定案核销状态查询 API返回值
// tmall.msf.identify.status.query
//
// 喵师傅定案核销状态查询，供服务商erp系统调用
type TmallMsfIdentifyStatusQueryAPIResponse struct {
	model.CommonResponse
	TmallMsfIdentifyStatusQueryAPIResponseModel
}

// TmallMsfIdentifyStatusQueryAPIResponseModel is 喵师傅定案核销状态查询 成功返回结果
type TmallMsfIdentifyStatusQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_msf_identify_status_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果json串，其中identifyDate位核销日期，signTime为签到时间，identifyStatus位核销状态，1代表已经核销，identifyType为核销类型，取值为：0代表未核销,1代表10位核销码核销,2代表订单号核销,3代表手机号核销,4代表4位核销码核销,5代表通过poi核销
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}

package alsc

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmCardPagetmpAPIResponse 查询卡模板列表(支持数据下行) API返回值
// alibaba.alsc.crm.card.pagetmp
//
// 查询卡模板列表(支持数据下行)
// 当传递了数据下行参数:
//      * isDeleted,lastMaxId,gmtModified,num时,进行数据下行处理,返回结果不带分页信息
//      * 否则分页查询卡模板,返回结果带有分页信息
type AlibabaAlscCrmCardPagetmpAPIResponse struct {
	model.CommonResponse
	AlibabaAlscCrmCardPagetmpAPIResponseModel
}

// AlibabaAlscCrmCardPagetmpAPIResponseModel is 查询卡模板列表(支持数据下行) 成功返回结果
type AlibabaAlscCrmCardPagetmpAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alsc_crm_card_pagetmp_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 分页返回模型
	Result *CommonPageResult `json:"result,omitempty" xml:"result,omitempty"`
}

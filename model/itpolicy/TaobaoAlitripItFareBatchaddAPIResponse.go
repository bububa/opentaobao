package itpolicy

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitripitfarebatchaddAPIResponse 【国际机票自有政策】批量添加 API返回值
// taobao.alitrip.it.fare.batchadd
//
// 支持自有政策和销售规则批量添加，支持携程的数据格式。淘宝格式为list [object] to json string，object的属性和单条接口一致。每个接入方最多同时只能有1个处理中的导入任务，超过后直接返回失败。文件一定要zip压缩，压缩后大小不超过5M，编码格式utf-8
type TaobaoalitripitfarebatchaddAPIResponse struct {
	model.CommonResponse
	TaobaoalitripitfarebatchaddAPIResponseModel
}

// TaobaoalitripitfarebatchaddAPIResponseModel is 【国际机票自有政策】批量添加 成功返回结果
type TaobaoalitripitfarebatchaddAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_it_fare_batchadd_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// json格式的字符串，扩展属性，预留
	ExtendAttributes string `json:"extend_attributes,omitempty" xml:"extend_attributes,omitempty"`
	// 任务id，可以根据任务id调用querytask查询执行结果
	TeskId int64 `json:"tesk_id,omitempty" xml:"tesk_id,omitempty"`
}

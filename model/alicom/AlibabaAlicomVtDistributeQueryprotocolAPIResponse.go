package alicom

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalicomvtdistributequeryprotocolAPIResponse 通信业务外放协议查询 API返回值
// alibaba.alicom.vt.distribute.queryprotocol
//
// 通信业务外放协议查询
type AlibabaalicomvtdistributequeryprotocolAPIResponse struct {
	model.CommonResponse
	AlibabaalicomvtdistributequeryprotocolAPIResponseModel
}

// AlibabaalicomvtdistributequeryprotocolAPIResponseModel is 通信业务外放协议查询 成功返回结果
type AlibabaalicomvtdistributequeryprotocolAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alicom_vt_distribute_queryprotocol_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回对象
	Result *CommonResult `json:"result,omitempty" xml:"result,omitempty"`
}

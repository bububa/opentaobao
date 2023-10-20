package ascpchannel

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaascpchanneldistributorproductlistAPIResponse 供应链渠道中心淘外分销品批量查询(分销商专用) API返回值
// alibaba.ascp.channel.distributor.product.list
//
// 此api为淘外分销的品批量查询标准api，淘外分销商专用
type AlibabaascpchanneldistributorproductlistAPIResponse struct {
	model.CommonResponse
	AlibabaascpchanneldistributorproductlistAPIResponseModel
}

// AlibabaascpchanneldistributorproductlistAPIResponseModel is 供应链渠道中心淘外分销品批量查询(分销商专用) 成功返回结果
type AlibabaascpchanneldistributorproductlistAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ascp_channel_distributor_product_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值包装,result为返回具体消息内容
	ProductListResponse *ResultWrapper `json:"product_list_response,omitempty" xml:"product_list_response,omitempty"`
}

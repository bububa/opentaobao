package btrip

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商旅成本中心转换为外部成本中心 API返回值 
alitrip.btrip.cost.center.transfer

商旅成本中心转换为外部成本中心
*/
type AlitripBtripCostCenterTransferAPIResponse struct {
    model.CommonResponse
    AlitripBtripCostCenterTransferAPIResponseModel
}

// 商旅成本中心转换为外部成本中心 成功返回结果
type AlitripBtripCostCenterTransferAPIResponseModel struct {
    XMLName xml.Name `xml:"alitrip_btrip_cost_center_transfer_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回对象
    Result   *BcmcResult `json:"result,omitempty" xml:"result,omitempty"`
}

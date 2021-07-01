package btrip

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取火车票订单列表 API返回值 
alitrip.btrip.train.order.search

第三方OA厂商获取自己的火车票数据
*/
type AlitripBtripTrainOrderSearchAPIResponse struct {
    model.CommonResponse
    AlitripBtripTrainOrderSearchAPIResponseModel
}

// 获取火车票订单列表 成功返回结果
type AlitripBtripTrainOrderSearchAPIResponseModel struct {
    XMLName xml.Name `xml:"alitrip_btrip_train_order_search_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回结果
    Result   *BtriphomeResultSupport `json:"result,omitempty" xml:"result,omitempty"`
}

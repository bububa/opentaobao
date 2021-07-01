package xhotelitem

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
复杂价格推送接口（全量更新） API返回值 
taobao.xhotel.multiplerate.update

酒店产品库复杂rate（多人价，连住价等）更新
同时完全涵盖taobao.xhotel.rate.update的功能
*/
type TaobaoXhotelMultiplerateUpdateAPIResponse struct {
    model.CommonResponse
    TaobaoXhotelMultiplerateUpdateAPIResponseModel
}

// 复杂价格推送接口（全量更新） 成功返回结果
type TaobaoXhotelMultiplerateUpdateAPIResponseModel struct {
    XMLName xml.Name `xml:"xhotel_multiplerate_update_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // gid-rpid-occupancy-lengthofstay<br/>商品ID-房价ID-入住人数-连住天数
    GidAndRpidOccupancyLengthofstay   string `json:"gid_and_rpid_occupancy_lengthofstay,omitempty" xml:"gid_and_rpid_occupancy_lengthofstay,omitempty"`
}

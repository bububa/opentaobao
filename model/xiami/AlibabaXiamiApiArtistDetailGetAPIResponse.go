package xiami

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
艺人详情 API返回值 
alibaba.xiami.api.artist.detail.get

艺人详情
*/
type AlibabaXiamiApiArtistDetailGetAPIResponse struct {
    model.CommonResponse
    AlibabaXiamiApiArtistDetailGetAPIResponseModel
}

// 艺人详情 成功返回结果
type AlibabaXiamiApiArtistDetailGetAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_xiami_api_artist_detail_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回数据
    Data   *StandardArtist `json:"data,omitempty" xml:"data,omitempty"`
}

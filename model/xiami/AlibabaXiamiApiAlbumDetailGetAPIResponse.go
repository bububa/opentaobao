package xiami

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaXiamiApiAlbumDetailGetAPIResponse
虾米音乐专辑详情接口 API返回值
alibaba.xiami.api.album.detail.get

虾米音乐专辑详情接口 */
type AlibabaXiamiApiAlbumDetailGetAPIResponse struct {
	model.CommonResponse
	AlibabaXiamiApiAlbumDetailGetAPIResponseModel
}

// AlibabaXiamiApiAlbumDetailGetAPIResponseModel is 虾米音乐专辑详情接口 成功返回结果
type AlibabaXiamiApiAlbumDetailGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_xiami_api_album_detail_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 专辑资料
	Data *AlbumDetail `json:"data,omitempty" xml:"data,omitempty"`
}

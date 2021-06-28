package xiami

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
虾米音乐专辑详情接口 APIResponse
alibaba.xiami.api.album.detail.get

虾米音乐专辑详情接口
*/
type AlibabaXiamiApiAlbumDetailGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_xiami_api_album_detail_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 专辑资料
    
    Data   *AlbumDetail `json:"data,omitempty" xml:"
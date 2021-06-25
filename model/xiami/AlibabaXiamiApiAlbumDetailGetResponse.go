package xiami

import (
    "github.com/bububa/opentaobao/model"
)

/* 
虾米音乐专辑详情接口 APIResponse
alibaba.xiami.api.album.detail.get

虾米音乐专辑详情接口
*/
type AlibabaXiamiApiAlbumDetailGetAPIResponse struct {
    model.CommonResponse
    Response *AlibabaXiamiApiAlbumDetailGetResponse `json:"alibaba_xiami_api_album_detail_get_response,omitempty"`
}

type AlibabaXiamiApiAlbumDetailGetResponse struct {

    // 专辑资料
    Data   *AlbumDetail `json:"data,omitempty"`

}

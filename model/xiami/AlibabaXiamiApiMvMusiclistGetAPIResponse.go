package xiami

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaXiamiApiMvMusiclistGetAPIResponse
乐馆mv列表 API返回值
alibaba.xiami.api.mv.musiclist.get

乐馆mv列表 */
type AlibabaXiamiApiMvMusiclistGetAPIResponse struct {
	model.CommonResponse
	AlibabaXiamiApiMvMusiclistGetAPIResponseModel
}

// AlibabaXiamiApiMvMusiclistGetAPIResponseModel is 乐馆mv列表 成功返回结果
type AlibabaXiamiApiMvMusiclistGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_xiami_api_mv_musiclist_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// mv_music_list_result
	MvMusicListResult *AlibabaXiamiApiMvMusiclistGetStruct `json:"mv_music_list_result,omitempty" xml:"mv_music_list_result,omitempty"`
}

package cloudgame

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCgameContentDistributionFileDownloadUpdateAPIResponse 文件下载回调 API返回值
// alibaba.cgame.content.distribution.file.download.update
//
// 提供内容服务器更新文件下载状态的能力
type AlibabaCgameContentDistributionFileDownloadUpdateAPIResponse struct {
	model.CommonResponse
	AlibabaCgameContentDistributionFileDownloadUpdateAPIResponseModel
}

// AlibabaCgameContentDistributionFileDownloadUpdateAPIResponseModel is 文件下载回调 成功返回结果
type AlibabaCgameContentDistributionFileDownloadUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_cgame_content_distribution_file_download_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 文件下载是否成功
	Succeeded bool `json:"succeeded,omitempty" xml:"succeeded,omitempty"`
}

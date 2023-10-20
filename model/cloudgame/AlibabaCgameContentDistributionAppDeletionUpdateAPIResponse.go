package cloudgame

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCgameContentDistributionAppDeletionUpdateAPIResponse 游戏删除回调 API返回值
// alibaba.cgame.content.distribution.app.deletion.update
//
// 游戏删除回调
type AlibabaCgameContentDistributionAppDeletionUpdateAPIResponse struct {
	model.CommonResponse
	AlibabaCgameContentDistributionAppDeletionUpdateAPIResponseModel
}

// AlibabaCgameContentDistributionAppDeletionUpdateAPIResponseModel is 游戏删除回调 成功返回结果
type AlibabaCgameContentDistributionAppDeletionUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_cgame_content_distribution_app_deletion_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 游戏是否成功
	Succeeded bool `json:"succeeded,omitempty" xml:"succeeded,omitempty"`
}

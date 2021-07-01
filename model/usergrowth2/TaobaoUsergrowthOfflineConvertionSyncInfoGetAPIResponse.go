package usergrowth2

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoUsergrowthOfflineConvertionSyncInfoGetAPIResponse
查询 用增线下业务  转化数据是否同步完成 API返回值
taobao.usergrowth.offline.convertion.sync.info.get

为手淘线下合作的渠道，提供对外查询数据是否更新完毕接口 */
type TaobaoUsergrowthOfflineConvertionSyncInfoGetAPIResponse struct {
	model.CommonResponse
	TaobaoUsergrowthOfflineConvertionSyncInfoGetAPIResponseModel
}

// TaobaoUsergrowthOfflineConvertionSyncInfoGetAPIResponseModel is 查询 用增线下业务  转化数据是否同步完成 成功返回结果
type TaobaoUsergrowthOfflineConvertionSyncInfoGetAPIResponseModel struct {
	XMLName xml.Name `xml:"usergrowth_offline_convertion_sync_info_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TaobaoUsergrowthOfflineConvertionSyncInfoGetResult `json:"result,omitempty" xml:"result,omitempty"`
}

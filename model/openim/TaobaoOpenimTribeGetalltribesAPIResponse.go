package openim

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOpenimTribeGetalltribesAPIResponse
获取用户群列表 API返回值
taobao.openim.tribe.getalltribes

OPENIM群服务获取用户群列表 */
type TaobaoOpenimTribeGetalltribesAPIResponse struct {
	model.CommonResponse
	TaobaoOpenimTribeGetalltribesAPIResponseModel
}

// TaobaoOpenimTribeGetalltribesAPIResponseModel is 获取用户群列表 成功返回结果
type TaobaoOpenimTribeGetalltribesAPIResponseModel struct {
	XMLName xml.Name `xml:"openim_tribe_getalltribes_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 群列表信息
	TribeInfoList []TribeInfo `json:"tribe_info_list,omitempty" xml:"tribe_info_list>tribe_info,omitempty"`
}
